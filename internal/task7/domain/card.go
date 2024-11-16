package domain

import (
	"fmt"
	"time"
)

type Card struct {
	Id         int64
	AccountId  int64
	CardNumber string `gorm:"column:cnumber"`
	CVV        string
	CreatedAt  time.Time
	ExpiryDate time.Time `gorm:"column:expired_at"`
	Status     string    `gorm:"column:cstatus"`
}

func (Card) TableName() string {
	return "card"
}

func (Card) Head() []string {
	return []string{"ID", "Account ID", "Card Number", "CVV", "Created At", "Expiry Date", "Status"}
}

func (c Card) Row() []string {
	return []string{
		fmt.Sprintf("%d", c.Id),
		fmt.Sprintf("%d", c.AccountId),
		c.CardNumber,
		c.CVV,
		c.CreatedAt.Format(time.DateOnly),
		c.ExpiryDate.Format(time.DateOnly),
		c.Status,
	}
}
