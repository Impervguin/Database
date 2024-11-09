package task6

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

// 7. Вызвать хранимую процедуру (написанную в третьей лабораторной работе);

type CallPromoRandomResponse struct {
	CountBefore int
	CountAfter  int
}

func (r *CallPromoRandomResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Active accounts before", "Active accounts after"})
	t.AppendRow(table.Row{fmt.Sprintf("%d", r.CountBefore), fmt.Sprintf("%d", r.CountAfter)})
	return t.Render()
}

const getActiveAccountsCount = `
SELECT COUNT(*) FROM loan WHERE lstatus = 'active'
`

const callPromoRandom = `
CALL PromoRandom()
`

func (t6s *Task6Storage) CallPromoRandom() (*CallPromoRandomResponse, error) {
	row := t6s.conn.QueryRow(getActiveAccountsCount)

	var before int
	err := row.Scan(&before)
	if err != nil {
		return nil, err
	}

	_, err = t6s.conn.Exec(callPromoRandom)
	if err != nil {
		return nil, err
	}

	row = t6s.conn.QueryRow(getActiveAccountsCount)
	var after int
	err = row.Scan(&after)
	if err != nil {
		return nil, err
	}
	return &CallPromoRandomResponse{
		CountBefore: before,
		CountAfter:  after,
	}, nil
}
