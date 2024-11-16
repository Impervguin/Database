package task7

import (
	"DatabaseCourse/internal/task7/domain"
	"errors"
	"time"

	"gorm.io/gorm"
)

func (t7s *Task7Storage) CreateCard(card *domain.Card) error {
	card.CreatedAt = time.Now()
	card.ExpiryDate = time.Now().Add(time.Hour * 24 * 30 * 12 * 3)
	card.Status = "active"
	res := t7s.db.Omit("id").Create(&card)
	return res.Error
}

func (t7s *Task7Storage) BlockCard(cardId int64) (*domain.Card, error) {
	card := &domain.Card{}
	res := t7s.db.Where("id =?", cardId).First(card)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, ErrResultEmpty
		}
		return nil, res.Error
	}
	card.Status = "blocked"
	res = t7s.db.Save(card)
	return card, res.Error
}

func (t7s *Task7Storage) DeleteCard(cardId int64) (*domain.Card, error) {
	card := &domain.Card{}
	res := t7s.db.Where("id =?", cardId).First(card)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, ErrResultEmpty
		}
		return nil, res.Error
	}
	res = t7s.db.Delete(card)
	return card, res.Error
}
