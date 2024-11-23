package task9

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx"
)

const (
	maxConn        = 10
	acquireTimeout = time.Minute
)

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
}

func NewPgStorage(ctx context.Context) (*PgStorage, error) {
	conf, err := pgx.ParseURI(DSN())

	if err != nil {
		return nil, err
	}
	poolConf := &pgx.ConnPoolConfig{
		ConnConfig:     conf,
		MaxConnections: maxConn,
		AcquireTimeout: acquireTimeout,
	}
	conn, err := pgx.NewConnPool(*poolConf)
	if err != nil {
		return nil, err
	}

	return &PgStorage{conn: conn}, nil
}

type PgStorage struct {
	conn *pgx.ConnPool
}

const getWealthyClients = `
SELECT client.id, first_name, last_name, phone_number, email, address, SUM(balance) as sb
FROM client JOIN account ON client.id = account.client_id
WHERE account.atype != 'credit' AND account.astatus = 'active'
GROUP BY client.id, first_name, last_name, phone_number, email
ORDER BY SUM(balance) DESC LIMIT $1
`

func (pgs *PgStorage) GetWealthyClients(limit int) ([]WealthyClient, error) {
	rows, err := pgs.conn.Query(getWealthyClients, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := make([]WealthyClient, 0, limit)
	for rows.Next() {
		var client WealthyClient
		err := rows.Scan(&client.Id, &client.FirstName, &client.LastName, &client.PhoneNumber, &client.Email, &client.Address, &client.Balance)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

const createAccount = `
INSERT INTO account (client_id, balance, interest, atype, astatus, created_at)
VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id;
`

func (pgs *PgStorage) CreateAccount(acc *Account) error {
	acc.AccountStatus = "active"
	err := pgs.conn.QueryRow(createAccount, acc.ClientId, acc.Balance, acc.Interest, acc.AccountType, acc.AccountStatus).Scan(&acc.Id)
	if err != nil {
		return err
	}
	return nil
}

const deleteRandomAccount = `
DELETE FROM account WHERE account.id = (SELECT account.id FROM Account WHERE account.id NOT IN (SELECT account_id FROM card) AND account.id NOT IN (SELECT account_id FROM transaction) ORDER BY RANDOM() LIMIT 1)
`

func (pgs *PgStorage) DeleteRandomAccount() error {
	_, err := pgs.conn.Exec(deleteRandomAccount)
	return err
}

const updateAccount = `
UPDATE account SET
client_id = $2,
balance = $3, 
interest = $4, 
atype = $5, 
astatus = $6, 
created_at = $7
WHERE id = $1
`

func (pgs *PgStorage) UpdateAccount(account *Account) error {
	_, err := pgs.conn.Exec(updateAccount, account.Id, account.ClientId, account.Balance, account.Interest, account.AccountType, account.AccountStatus, account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

const getRandomActiveAccount = `
SELECT id, client_id, balance, interest, atype, astatus, created_at
FROM account
WHERE astatus = 'active'
ORDER BY RANDOM()
LIMIT 1
`

func (pgs *PgStorage) GetRandomActiveAccount() (*Account, error) {
	row := pgs.conn.QueryRow(getRandomActiveAccount)

	var acc Account
	err := row.Scan(&acc.Id, &acc.ClientId, &acc.Balance, &acc.Interest, &acc.AccountType, &acc.AccountStatus, &acc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

const getRandomClientId = `SELECT client.id FROM client ORDER BY RANDOM() LIMIT 1`

func (pgs *PgStorage) GetRandomClientId() (int64, error) {
	var clientId int64
	err := pgs.conn.QueryRow(getRandomClientId).Scan(&clientId)
	if err != nil {
		return 0, err
	}
	return clientId, nil
}
