package task7

import (
	"DatabaseCourse/internal/task7/domain"
	"fmt"

	"gorm.io/gorm"
)

// Возвращает суммарный баланс человека на его счетах, без учёта кредитных

// Вызов функции

func (t7s *Task7Storage) GetClientTotalBalance(clientId int64) (*domain.ClientBalance, error) {
	var cb domain.ClientBalance
	res := t7s.db.Select("? as id, TotalBalance(?) as total_balance", fmt.Sprintf("%d", clientId), fmt.Sprintf("%d", clientId)).First(&cb)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return nil, ErrResultEmpty
		}
		return nil, res.Error
	}
	return &cb, nil
}
