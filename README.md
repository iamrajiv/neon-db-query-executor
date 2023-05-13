<div align="center">
<img src="assets/neon-db-query-executor.svg" height="auto" width="400" />
<br />
<h1>Neon DB Query Executor</h1>
<p>
Simple Go application to execute SQL queries against a Neon database
</p>
<a href="https://github.com/iamrajiv/neon-db-query-executor/network/members"><img src="https://img.shields.io/github/forks/iamrajiv/neon-db-query-executor?color=0969da&style=for-the-badge" height="auto" width="auto" /></a>
<a href="https://github.com/iamrajiv/neon-db-query-executor/stargazers"><img src="https://img.shields.io/github/stars/iamrajiv/neon-db-query-executor?color=0969da&style=for-the-badge" height="auto" width="auto" /></a>
<a href="https://github.com/iamrajiv/neon-db-query-executor/blob/main/LICENSE"><img src="https://img.shields.io/github/license/iamrajiv/neon-db-query-executor?color=0969da&style=for-the-badge" height="auto" width="auto" /></a>
</div>

## About

The SQL Query Executor for Neon Database is a lightweight Go-based application that allows you to execute SQL queries stored in files against a Neon database. It provides a simple and efficient way to run SQL queries and print the results.

[Neon](https://neon.tech/) is a fully managed serverless PostgreSQL. Neon separates storage and computing to offer modern developer features such as serverless, branching, bottomless storage, and more. Learn more about Neon [here](https://neon.tech/docs/introduction).

The Neon DB Query Executor has the following features:

- **Query Execution**
  - Execute SQL queries stored in files against a Neon database.
  - Support for executing multiple queries in a single file.
- **Database Connectivity**
  - Connect to a Neon database using Go's **`database/sql`** package.
  - Configuration of database connection details via environment variables.
  - Secure connection establishment with SSL/TLS support (**`sslmode=verify-full`**).
- **Comment Handling**
  - Handle comments in SQL files for additional information or explanations.
  - Support for both single-line comments (**`- Comment`**) and multi-line comments (**`/* Comment */`**).
  - Remove comments from the SQL file before query execution.
- **Skip Command Exclusion**
  - Exclude specific SQL commands from being displayed in the results.
  - Skip commands such as **`CREATE`**, **`ALTER`**, **`DROP`**, **`INSERT`**, etc., that do not produce output results.
  - Enhance result clarity by focusing on queries that generate result sets.
- **Result Presentation**
  - Print query results in a tabular format for easy viewing and analysis.
  - Utilize the **`tablewriter`** package to render formatted tables.
  - Display column names as table headers for clear representation.
- **Execution Time Tracking**
  - Calculate and display the execution time for each query to measure performance.
  - Track the elapsed time from query execution start to completion.
  - Print the execution time in seconds with precision.
- **Total Execution Time**
  - Show the total execution time for all queries in the file.
  - Accumulate the execution times of individual queries to calculate the total time taken.

These features collectively enable the SQL Query Executor for Neon Database to provide a comprehensive solution for executing SQL queries against a Neon database. It offers flexibility in query execution, result presentation, comment handling, and execution time tracking. The skip command exclusion feature further enhances output clarity by excluding non-result-producing commands.

The folder structure of the project is as follows:

```shell
.
├── .env
├── LICENSE
├── README.md
├── assets
│   └── neon-db-query-executor.svg
├── go.mod
├── go.sum
├── main.go
└── queries.sql
```

## Usage

1. To set up a project in Neon, follow the instructions [here](https://neon.tech/docs/get-started-with-neon/setting-up-a-project).
2. Configure the database connection details in a `.env` file.
3. Write your SQL queries in `queries.sql` file.
4. Run the application using `go run main.go`.

#### Instructions for writing environment variables

To run the project locally, create a `.env` file and add the following environment variables:

- **`DB_USER`**: Username for the Neon database.
- **`DB_PASSWORD`**: Password for the Neon database.
- **`DB_NAME`**: Name of the Neon database. When creating a project, you can choose any name for the database, so make sure to update this variable accordingly.
- **`DB_ENDPOINT_ID`**: ID of the Neon database endpoint. The endpoint ID will change for every new branch you create.
- **`DB_PROXY_HOST`**: Host for the Neon database proxy. When creating a new project, you will be asked to choose a region from a dropdown list of available regions.
- **`DB_HOST`**: Leave this as it is. It will be automatically set based on the values of **`DB_ENDPOINT_ID`** and **`DB_PROXY_HOST`**. #**`verify-full`**
- **`DB_SSLMODE`**: Set the SSL mode for the database connection. By default, it is set to `verify-full`.

To get the values for the above environment variables, follow the instructions [here](https://neon.tech/docs/get-started-with-neon/setting-up-a-project#step-3-connect-to-your-neon-database).

Once the Neon project is successfully created, connection details are generated for accessing the default `neondb` database. These connection details can be saved or retrieved later from the connection details widget on the Neon dashboard. They provide the necessary environment variables for connecting to the database.

For direct connection, the connection details look like this:

```shell
psql 'postgresql://<DB_USER>:<DB_PASSWORD>@<DB_HOST>:/<DB_NAME>'
```

For example:

```shell
psql 'postgresql://test:123@ep-test.us-east-2.aws.neon.tech/neondb'
```

So, here `DB_USER` is `test`, `DB_PASSWORD` is `123`, `DB_ENDPOINT_ID` is `ep-test`, `DB_PROXY_HOST` is `us-east-2.aws.neon.tech`, and `DB_NAME` is `neondb`.

Some applications open numerous connections, with most eventually becoming inactive. This behavior can often be attributed to database driver limitations, to running many instances of an application, or to applications with serverless functions. With regular PostgreSQL, new connections are rejected when reaching the `max_connections` limit. To overcome this limitation, Neon supports connection pooling using [PgBouncer](https://www.pgbouncer.org/), allowing Neon to support up to 10000 concurrent connections.

Enabling connection pooling in Neon requires adding a `-pooler` suffix to the compute endpoint ID, which is part of the hostname. Connection requests that specify the `-pooler` suffix use a pooled connection.

Add the `-pooler` suffix to the endpoint ID, as shown:

```shell
psql 'postgresql://test:123@ep-test-pooler.us-east-2.aws.neon.tech/neondb'
```

So, here `DB_USER` is `test`, `DB_PASSWORD` is `123`, `DB_ENDPOINT_ID` is `ep-test-pooler`, `DB_PROXY_HOST` is `us-east-2.aws.neon.tech`, and `DB_NAME` is `neondb`.

#### Instructions for writing SQL queries

1. All SQL queries should be written in the **`queries.sql`** file.
2. The **`queries.sql`** file should be in the same directory as the **`main.go`** file.
3. Each SQL query should be on a separate line and should end with a semicolon **`;`**.
4. Make sure there are no empty lines or extra spaces before or after each SQL query.
5. The Go program can handle SQL queries with empty lines and white spaces, single-line comments, and multi-line comments, but it's best to write queries in a proper format without unnecessary empty lines or extra spaces.

#### Demonstration

Before the demonstration, I set up the new project and configured the database connection details in a `.env` file.

I wrote sample queries in the `queries.sql` file. Something like this:

```sql
-- Create a table
CREATE TABLE playing_with_neon (id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL);

-- Insert some data
INSERT INTO playing_with_neon (name, value) SELECT 'Data ' || generate_series, random() FROM generate_series(1, 10);

-- Query the table
SELECT * FROM playing_with_neon;

-- Drop the table
DROP TABLE playing_with_neon;

```

When I execute the `go run main.go` command to run the application, I receive the following output:

```shell
➜  neon-db-query-executor git:(main) ✗ go run main.go
Query: -- Create a table
CREATE TABLE playing_with_neon (id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL)
Time Taken: 1.836177 seconds

Query: -- Insert some data
INSERT INTO playing_with_neon (name, value) SELECT 'Data ' || generate_series, random() FROM generate_series(1, 10)
Time Taken: 0.251565 seconds

+--------------+--------------+--------------+
|      ID      |     NAME     |    VALUE     |
+--------------+--------------+--------------+
| 0xc000015590 | 0xc0000155b0 | 0xc0000155c0 |
| 0xc0000155f0 | 0xc000015610 | 0xc000015620 |
| 0xc000015650 | 0xc000015670 | 0xc000015680 |
| 0xc0000156b0 | 0xc0000156d0 | 0xc0000156e0 |
| 0xc000015710 | 0xc000015730 | 0xc000015740 |
| 0xc000015770 | 0xc000015790 | 0xc0000157a0 |
| 0xc0000157d0 | 0xc0000157f0 | 0xc000015800 |
| 0xc000015830 | 0xc000015850 | 0xc000015860 |
| 0xc000015890 | 0xc0000158b0 | 0xc0000158c0 |
+--------------+--------------+--------------+
Query: -- Query the table
SELECT * FROM playing_with_neon
Time Taken: 0.244171 seconds

Query: -- Drop the table
DROP TABLE playing_with_neon
Time Taken: 0.244211 seconds

Total Time Taken: 2.576124 seconds
```

## License

[MIT](https://github.com/iamrajiv/neon-db-query-executor/blob/main/LICENSE)
