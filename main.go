package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/olekukonko/tablewriter"
)

// Connect to the database using the environment variables and return a pointer to the *sql.DB object
func connectDB() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"), os.Getenv("DB_SSLMODE"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Execute the SQL query on the database and return the result rows and column names
func executeQuery(db *sql.DB, query string) (*sql.Rows, []string, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, nil, err
	}

	return rows, columns, nil
}

// Print the result rows and column names in a table using the "tablewriter" package
func printResults(rows *sql.Rows, columns []string) {
	data := [][]string{}
	for rows.Next() {
		var row []interface{}
		for range columns {
			var value interface{}
			row = append(row, &value)
		}
		err := rows.Scan(row...)
		if err != nil {
			log.Fatal(err)
		}

		rowValues := []string{}
		for _, value := range row {
			if value == nil {
				rowValues = append(rowValues, "NULL")
			} else {
				rowValues = append(rowValues, fmt.Sprintf("%v", value))
			}
		}
		data = append(data, rowValues)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(columns)
	table.AppendBulk(data)
	table.Render()
}

// Print the time taken to execute the SQL query
func printQueryTime(query string, duration time.Duration) {
	fmt.Printf("Query: %s\nTime Taken: %f seconds\n\n", query, duration.Seconds())
}

// Print the total time taken to execute all SQL queries
func printTotalTime(totalTime time.Duration) {
	fmt.Printf("Total Time Taken: %f seconds\n", totalTime.Seconds())
}

// Check if the SQL query is a command that does not produce output
func isSkipCommand(query string) bool {
	// List of SQL commands that do not produce output
	skipCommands := []string{
		"CREATE", "ALTER", "INSERT", "DROP", "UPDATE", "DELETE",
		"SET", "GRANT", "REVOKE", "COMMIT", "ROLLBACK",
	}

	for _, command := range skipCommands {
		if strings.HasPrefix(strings.ToUpper(query), command) {
			return true
		}
	}
	return false
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to the database
	db := connectDB()
	defer db.Close()

	// Read SQL queries from a file
	queryBytes, err := os.ReadFile("queries.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Remove comments from the SQL file
	re := regexp.MustCompile(`--.*$|/\*[\s\S]*?\*/`)
	cleanQueryBytes := re.ReplaceAll(queryBytes, []byte(""))

	// Split the SQL queries by semicolons
	queries := strings.Split(string(cleanQueryBytes), ";")

	// Execute each SQL query and print the results
	totalTime := time.Duration(0)
	for _, query := range queries {
		trimmedQuery := strings.TrimSpace(query)
		if trimmedQuery == "" {
			continue
		}

		// Measure the time taken to execute the query
		startTime := time.Now()

		// Execute the SQL query and get the result rows and column names
		rows, columns, err := executeQuery(db, trimmedQuery)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// Print the result rows and column names if the query is not a command that does not produce output
		if rows.Next() || isSkipCommand(trimmedQuery) {
			printResults(rows, columns)
		}

		// Print the time taken to execute the query
		duration := time.Since(startTime)
		totalTime += duration
		printQueryTime(trimmedQuery, duration)
	}

	// Print the total time taken to execute all queries
	printTotalTime(totalTime)
}
