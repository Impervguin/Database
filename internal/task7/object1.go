package task7

import "DatabaseCourse/internal/task7/domain"

// Получение всех клиентов по алфавитному порядку

func (t7s *Task7Storage) GetClientsAlphabetically() ([]domain.Client, error) {
	clients := make([]domain.Client, 0)
	res := t7s.db.Order("last_name ASC").Order("first_name ASC").Find(&clients)
	if res.Error != nil {
		return nil, res.Error
	}
	return clients, nil
}
