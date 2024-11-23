package task9

import (
	"encoding/json"
	"time"

	"golang.org/x/exp/rand"
)

type WealthyClient struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	Balance     float64
}

func (client *WealthyClient) MarshalBinary() ([]byte, error) {
	return json.Marshal(client)
}

func (client *WealthyClient) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, client)
}

type Account struct {
	Id            int64
	ClientId      int64
	Balance       float64
	Interest      float64
	AccountType   string
	AccountStatus string
	CreatedAt     time.Time
}

func CreateRandonAccount(clientId int64) (*Account, error) {
	return &Account{
		ClientId:      clientId,
		Balance:       rand.Float64() * 10000,
		Interest:      rand.Float64() * 10,
		AccountType:   "checking",
		AccountStatus: "active",
		CreatedAt:     time.Now(),
	}, nil
}

func (acc *Account) ApplyInterest() {
	acc.Balance += acc.Balance * (acc.Interest / 100)
}
