package task7

import (
	"DatabaseCourse/internal/task7/domain"
	"errors"

	"gorm.io/gorm"
)

// Однотабличный запрос на выборку

func (t *Task7Storage) GetCardById(id int64) (*domain.Card, error) {
	card := &domain.Card{}
	res := t.db.Where("id = ?", id).First(card)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, ErrResultEmpty
		}
		return nil, res.Error
	}
	return card, nil
}
