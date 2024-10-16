-- Хранимая процедура доступа к метаданным

-- SELECT * FROM information_schema.columns WHERE table_name = 'client';

DROP TABLE IF EXISTS my_columns;

CREATE TABLE IF NOT EXISTS my_columns (
    column_name TEXT,
    data_type TEXT
);

CREATE OR REPLACE PROCEDURE PrintTableAttributes(tablename text)
LANGUAGE PLPGSQL
AS $$
BEGIN
    INSERT INTO my_columns
    SELECT column_name, data_type FROM information_schema.columns WHERE table_name = tablename;
END;
$$;

CALL PrintTableAttributes('client');
SELECT * FROM my_columns;