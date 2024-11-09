package task6

// 10. Выполнить вставку данных в созданную таблицу с использованием инструкции INSERT или COPY.

type Stocks struct {
	ID        int
	Company   string
	Price     float64
	Quantity  int
	Remaining int
	Dividend  float64
}

const insertStock = `
INSERT INTO stocks (company, price, quantity, remaining, dividend)
VALUES ($1, $2, $3, $3, $4)
`

func (t6s *Task6Storage) InsertStock(stock Stocks) error {
	_, err := t6s.conn.Exec(insertStock, stock.Company, stock.Price, stock.Quantity, stock.Dividend)
	return err
}
