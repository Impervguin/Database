package task7

import "DatabaseCourse/internal/task7/domain"

// Найти всех клиентов у которых есть хотя бы один займ

func (t7s *Task7Storage) GetClientsWithLoan() ([]domain.ShortClient, error) {
	clients := make([]domain.ShortClient, 0)
	res := t7s.db.
		// Select("", domain.ShortClient{}.Head()).
		Joins("INNER JOIN loan ON client.id = loan.client_id").
		Where("loan.lstatus = 'active'").
		Group("client.id, client.first_name, client.last_name, client.email, client.phone_number").
		Having("COUNT(loan.id) > 0").
		Find(&clients)
	if res.Error != nil {
		return nil, res.Error
	}
	return clients, nil
}
