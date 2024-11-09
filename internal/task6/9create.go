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

func (t6s *Task6Storage) CreateStockTable() error {
	_, err := t6s.conn.Exec(createStockTable)
	return err
}
