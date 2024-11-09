package task6

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

// 3. Выполнить запрос с ОТВ(CTE) и оконными функциями;

// Получить количество счетов по типам, мин, макс, средний баланс

type GetAccountTypesStatsResponse struct {
	AccountTypes []AccountTypeStats
}

func (r *GetAccountTypesStatsResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Account type", "Count", "Minimal balance", "Maximal balance", "Average balance"})
	t.AppendFooter(table.Row{"Account type", "Count", "Minimal balance", "Maximal balance", "Average balance"})
	for _, c := range r.AccountTypes {
		s := c.Row()
		row := make([]interface{}, 0, len(s))
		for _, v := range s {
			row = append(row, v)
		}
		t.AppendRow(row)
	}
	return t.Render()
}

type AccountTypeStats struct {
	AccountType    string
	Count          int
	MinBalance     float64
	MaxBalance     float64
	AverageBalance float64
}

func (s *AccountTypeStats) Row() []string {
	return []string{s.AccountType, fmt.Sprintf("%d", s.Count), fmt.Sprintf("%f", s.MinBalance), fmt.Sprintf("%f", s.MaxBalance), fmt.Sprintf("%f", s.AverageBalance)}
}

const getAccountTypesAndStats = `
WITH account_stats AS (
    SELECT atype, COUNT(*) AS count, MIN(balance) AS min_balance, MAX(balance) AS max_balance, AVG(balance) AS avg_balance
    FROM account
    GROUP BY atype
)
SELECT atype, count, min_balance, max_balance, avg_balance
FROM account_stats
`

func (t6s *Task6Storage) GetAccountTypesAndStats() (*GetAccountTypesStatsResponse, error) {
	rows, err := t6s.conn.Query(getAccountTypesAndStats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]AccountTypeStats, 0)
	for rows.Next() {
		var s AccountTypeStats
		err := rows.Scan(&s.AccountType, &s.Count, &s.MinBalance, &s.MaxBalance, &s.AverageBalance)
		if err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &GetAccountTypesStatsResponse{AccountTypes: stats}, nil
}
