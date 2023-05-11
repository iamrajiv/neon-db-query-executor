-- Create a table
CREATE TABLE playing_with_neon (id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL);

-- Insert some data
INSERT INTO playing_with_neon (name, value) SELECT 'Data ' || generate_series, random() FROM generate_series(1, 10);

-- Query the table
SELECT * FROM playing_with_neon;

-- Drop the table
DROP TABLE playing_with_neon;
