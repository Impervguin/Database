package task7

import "DatabaseCourse/internal/task7/domain"

// Найти клиентов у которых больше всего счетов(по количеству)

func (t7s *Task7Storage) GetClientsWithMostAccounts() ([]domain.AccountCount, error) {
	clientSums := make([]domain.AccountCount, 0)
	res := t7s.db.
		Select("client_id, first_name, last_name, phone_number, email, COUNT(*) as account_count").
		Joins("JOIN client ON client.id = account.client_id").
		Group("client_id, first_name, last_name, phone_number, email").
		Having("COUNT(*) = (?)", t7s.db.Select("MAX(c)").Table("(?)", t7s.db.Select("COUNT(*) as c").Table("account").Group("client_id"))).
		Find(&clientSums)

	if res.Error != nil {
		return nil, res.Error
	}
	return clientSums, nil
}
