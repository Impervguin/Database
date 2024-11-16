package domain

import "fmt"

type ClientSum struct {
	Id          int64 `gorm:"column:client_id"`
	AccountsSum float64
}

func (ClientSum) TableName() string {
	return "account"
}

func (ClientSum) Head() []string {
	return []string{"Client ID", "Accounts Sum"}
}

func (c ClientSum) Row() []string {
	return []string{fmt.Sprintf("%d", c.Id), fmt.Sprintf("%.2f", c.AccountsSum)}
}
