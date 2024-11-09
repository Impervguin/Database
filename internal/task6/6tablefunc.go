package task6

import "github.com/jedib0t/go-pretty/table"

// 6. Вызвать многооператорную или табличную функцию (написанную в третьей лабораторной работе);

// Вывести имена, телефоны и данные карточки клиентов, у которых они заблокированы

type GetBlockedClientsResponse struct {
	Clients []BlockedClient
}

func (r *GetBlockedClientsResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Name", "Phone number", "Card number", "Card status"})
	t.AppendFooter(table.Row{"Name", "Phone number", "Card number", "Card status"})
	for _, bc := range r.Clients {
		s := bc.Row()
		row := make([]interface{}, 0, len(s))
		for _, v := range s {
			row = append(row, v)
		}
		t.AppendRow(row)
	}
	return t.Render()
}

type BlockedClient struct {
	LFName      string
	PhoneNumber string
	CardNumber  string
	CardStatus  string
}

func (r *BlockedClient) Row() []string {
	return []string{r.LFName, r.PhoneNumber, r.CardNumber, r.CardStatus}
}

const getBlockedClients = `
SELECT bc.lf, bc.phone, bc.cnumber, bc.cstatus
FROM client
JOIN BlockedCards(client.id) AS bc ON TRUE
`

func (t6s *Task6Storage) GetBlockedClients() (*GetBlockedClientsResponse, error) {
	rows, err := t6s.conn.Query(getBlockedClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	blockedClients := make([]BlockedClient, 0)
	for rows.Next() {
		var bc BlockedClient
		err := rows.Scan(&bc.LFName, &bc.PhoneNumber, &bc.CardNumber, &bc.CardStatus)
		if err != nil {
			return nil, err
		}
		blockedClients = append(blockedClients, bc)
	}
	return &GetBlockedClientsResponse{Clients: blockedClients}, nil
}
