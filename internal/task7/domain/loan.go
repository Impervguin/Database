package domain

import (
	"fmt"
	"time"
)

type Loan struct {
	Id              int64
	ClientId        int64
	Amount          float64
	Interest        float64
	RemainingAmount float64
	MonthlyPayment  float64
	StartDate       time.Time
	EndDate         time.Time
	Status          string `gorm:"column:lstatuses"`
	Description     string `gorm:"column:ldescription"`
}

func (Loan) TableName() string {
	return "loan"
}

func (Loan) Head() []string {
	return []string{"ID", "Client ID", "Amount", "Interest", "Remaining Amount", "Monthly Payment", "Start Date", "End Date", "Status", "Description"}
}

func (l Loan) Row() []string {
	return []string{
		fmt.Sprintf("%d", l.Id),
		fmt.Sprintf("%d", l.ClientId),
		fmt.Sprintf("%f", l.Amount),
		fmt.Sprintf("%f", l.Interest),
		fmt.Sprintf("%f", l.RemainingAmount),
		fmt.Sprintf("%f", l.MonthlyPayment),
		l.StartDate.Format(time.RFC3339),
		l.EndDate.Format(time.RFC3339),
		l.Status,
		l.Description,
	}
}
