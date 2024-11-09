package task6

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

// 5. Вызвать скалярную функцию (написанную в третьей лабораторной работе)

// Возвращает суммарный баланс человека на его счетах, без учёта кредитных

type GetSumBalanceResponse struct {
	FirstName    string
	LastName     string
	TotalBalance float64
}

func (r *GetSumBalanceResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"First name", "Last Name", "Total balance"})
	t.AppendRow(table.Row{r.FirstName, r.LastName, fmt.Sprintf("%f", r.TotalBalance)})
	return t.Render()
}

const getSumBalance = `
SELECT first_name, last_name, TotalBalance($1)
FROM client
WHERE id = $1
`

func (t6s *Task6Storage) GetSumBalance(id int64) (*GetSumBalanceResponse, error) {
	row := t6s.conn.QueryRow(getSumBalance, id)

	var firstName, lastName string
	var totalBalance float64
	err := row.Scan(&firstName, &lastName, &totalBalance)
	if err != nil {
		return nil, err
	}

	return &GetSumBalanceResponse{
		FirstName:    firstName,
		LastName:     lastName,
		TotalBalance: totalBalance,
	}, nil
}
