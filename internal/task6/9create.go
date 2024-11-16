package task6

// 9. Создать таблицу в базе данных, соответствующую тематике БД

const createStockTable = `
CREATE TABLE IF NOT EXISTS stocks (
	id SERIAL PRIMARY KEY,
	company VARCHAR(255) NOT NULL UNIQUE,
	price DECIMAL(20, 2) NOT NULL CHECK(price >0),
    quantity INTEGER NOT NULL CHECK (quantity >0),
	remaining INTEGER NOT NULL CHECK (remaining>=0),
	dividend DECIMAL(20, 2) NOT NULL CHECK (dividend >=0)
)
`

const checkStockTable = `
SELECT EXISTS(
    SELECT *
    FROM information_schema.tables
    WHERE
      table_name = 'stocks'
);`

func (t6s *Task6Storage) CreateStockTable() error {
	_, err := t6s.conn.Exec(createStockTable)
	return err
}

func (t6s *Task6Storage) CheckStockTable() (bool, error) {
	var exists bool
	err := t6s.conn.QueryRow(checkStockTable).Scan(&exists)
	return exists, err
}
