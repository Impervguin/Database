package domain

import "fmt"

type AccountCount struct {
	Id           int64 `gorm:"column:client_id"`
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  string
	AccountCount uint
}

func (a AccountCount) TableName() string {
	return "account"
}

func (AccountCount) Head() []string {
	return []string{"Client ID", "Account Count"}
}

func (a AccountCount) Row() []string {
	return []string{fmt.Sprintf("%d", a.Id), fmt.Sprintf("%d", a.AccountCount)}
}
