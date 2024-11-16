package task7

import "DatabaseCourse/internal/task7/domain"

// Для каждого клиента найти сумму его счетов, кроме кредитных

func (t7s *Task7Storage) GetClientsAccountSums() ([]domain.ClientSum, error) {
	clients := make([]domain.ClientSum, 0)
	res := t7s.db.
		Select("client_id, SUM(balance) as accounts_sum").
		Where("atype != 'credit'").
		Group("client_id").
		Find(&clients)
	if res.Error != nil {
		return nil, res.Error
	}
	return clients, nil
}
