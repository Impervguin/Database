package task7

import "DatabaseCourse/internal/task7/domain"

// Найти все невыплаченные займы, оставшаяся сумма которых превышает 1 000 000, или ежемесячный платёж больше 10 000

func (t7s *Task7Storage) GetUnpaidLoansExceeding1Million() ([]domain.Loan, error) {
	loans := make([]domain.Loan, 0)
	res := t7s.db.Where("lstatus = 'active'").Where("remaining_amount > 1000000").Find(&loans)
	if res.Error != nil {
		return nil, res.Error
	}
	return loans, nil
}
