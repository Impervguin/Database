package domain

import (
	"fmt"
	"time"
)

type Client struct {
	Id          int64
	FirstName   string
	LastName    string
	dob         time.Time
	Email       string
	PhoneNumber string
	Address     string
	CreatedAt   time.Time
}

func (Client) TableName() string {
	return "client"
}

func (c Client) Row() []string {
	return []string{fmt.Sprintf("%d", c.Id), c.FirstName, c.LastName, c.Email, c.PhoneNumber, c.Address, c.CreatedAt.Format(time.RFC3339)}
}

func (Client) Head() []string {
	return []string{"ID", "First Name", "Last Name", "Email", "Phone Number", "Address", "Created At"}
}
