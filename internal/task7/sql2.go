package task7

import "DatabaseCourse/internal/task7/domain"

// Многотабличный запрос на выборку

func (t *Task7Storage) GetCardsForClient(clientId int64) ([]domain.Card, error) {
	cards := make([]domain.Card, 0)
	res := t.db.Joins("JOIN account on account.id = card.account_id").
		Where("client_id = ?", clientId).Find(&cards)
	if len(cards) == 0 {
		return nil, ErrResultEmpty
	}
	if res.Error != nil {
		return nil, res.Error
	}
	return cards, nil
}
