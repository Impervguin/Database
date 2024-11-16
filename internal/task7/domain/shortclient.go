package domain

import "fmt"

type ShortClient struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

func (ShortClient) TableName() string {
	return "client"
}

func (ShortClient) Head() []string {
	return []string{"ID", "First Name", "Last Name", "Email", "Phone Number"}
}

func (c ShortClient) Row() []string {
	return []string{fmt.Sprintf("%d", c.Id), c.FirstName, c.LastName, c.Email, c.PhoneNumber}
}
