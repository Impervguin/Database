package task6

import (
	"github.com/jedib0t/go-pretty/table"
)

// 8. Вызвать системную функцию или процедуру

type CurrentUserResponse struct {
	UserName string
}

func (r *CurrentUserResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Username"})
	t.AppendRow(table.Row{r.UserName})
	return t.Render()
}

const getCurrentUser = `
SELECT current_user
`

func (t6s *Task6Storage) GetCurrentUser() (*CurrentUserResponse, error) {
	row := t6s.conn.QueryRow(getCurrentUser)

	var userName string
	err := row.Scan(&userName)
	if err != nil {
		return nil, err
	}
	return &CurrentUserResponse{UserName: userName}, nil
}
