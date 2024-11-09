package task6

import "github.com/jedib0t/go-pretty/table"

// 1. Выполнить скалярный запрос

// Получить число всех клиентов в базе данных

type GetClientsCountResponse struct {
	ClientCount int
}

func (r *GetClientsCountResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Client count"})
	t.AppendRow(table.Row{r.ClientCount})
	return t.Render()
}

const getClientsCount = `
SELECT COUNT(*) FROM client 
`

func (t6s *Task6Storage) GetClientsCount() (*GetClientsCountResponse, error) {
	row := t6s.conn.QueryRow(getClientsCount)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &GetClientsCountResponse{ClientCount: count}, nil
}
