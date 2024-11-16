package domain

import "fmt"

type ClientBalance struct {
	Id           int64
	TotalBalance float64
}

func (ClientBalance) TableName() string {
	return "client"
}

func (ClientBalance) Head() []string {
	return []string{"Client ID", "Total Balance"}
}

func (c ClientBalance) Row() []string {
	return []string{fmt.Sprintf("%d", c.Id), fmt.Sprintf("%.2f", c.TotalBalance)}
}
