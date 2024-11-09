package task6

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
)

// 2. Выполнить запрос с несколькими соединениями (JOIN);

// Получить всех клиентов и номера их карт

type GetClientsAndCardsResponse struct {
	Clients []ClientAndCard
}

func (r *GetClientsAndCardsResponse) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"ID", "First name", "Last name", "Card Number", "Card type"})
	t.AppendFooter(table.Row{"ID", "First name", "Last name", "Card Number", "Card type"})
	for _, c := range r.Clients {
		s := c.Row()
		row := make([]interface{}, 0, len(s))
		for _, v := range s {
			row = append(row, v)
		}
		t.AppendRow(row)
	}
	return t.Render()
}

type ClientAndCard struct {
	ClientID         int
	ClientFirstName  string
	ClientSecondName string
	CardNumber       string
	CardType         string
}

func (r *ClientAndCard) Row() []string {
	return []string{fmt.Sprintf("%d", r.ClientID), r.ClientFirstName, r.ClientSecondName, r.CardNumber, r.CardType}
}

const getClientsAndCards = `
SELECT client.id, client.first_name, client.last_name, card.cnumber, account.atype
FROM client
JOIN account ON client.id = account.client_id
JOIN card ON account.id = card.account_id
`

func (t6s *Task6Storage) GetClientsAndCards() (*GetClientsAndCardsResponse, error) {
	rows, err := t6s.conn.Query(getClientsAndCards)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := make([]ClientAndCard, 0)
	for rows.Next() {
		var c ClientAndCard
		err := rows.Scan(&c.ClientID, &c.ClientFirstName, &c.ClientSecondName, &c.CardNumber, &c.CardType)
		if err != nil {
			return nil, err
		}
		clients = append(clients, c)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &GetClientsAndCardsResponse{Clients: clients}, nil
}
