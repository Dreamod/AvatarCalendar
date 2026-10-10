package day

import (
	"aiCalendarApp/pkg/db"
	"context"

	"gorm.io/gorm/clause"
)

type Repository struct {
	DataBase *db.Db
}

func NewRepository(dataBase *db.Db) *Repository {
	return &Repository{
		DataBase: dataBase,
	}
}

// Create создает запись в БД
func (repo *Repository) Create(ctx context.Context, day *Day) (*Day, error) {
	result := repo.DataBase.DB.WithContext(ctx).Create(day)
	if result.Error != nil {
		return nil, result.Error
	}
	return day, nil
}

// GetById получает запись по ID
func (repo *Repository) GetById(ctx context.Context, id string) (*Day, error) {
	var day Day
	result := repo.DataBase.DB.WithContext(ctx).First(&day, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &day, nil
}

// Update обновляет запись в БД
func (repo *Repository) Update(ctx context.Context, day *Day) (*Day, error) {
	result := repo.DataBase.DB.WithContext(ctx).Clauses(clause.Returning{}).Updates(day)
	if result.Error != nil {
		return nil, result.Error
	}
	return day, nil
}
