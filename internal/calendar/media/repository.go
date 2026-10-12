package media

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

func (repo *Repository) Create(ctx context.Context, media *Media) (*Media, error) {
	result := repo.DataBase.DB.WithContext(ctx).Create(media)
	if result.Error != nil {
		return nil, result.Error
	}
	return media, nil
}

// GetByDayId получает медиа файлы по ID дня
func (repo *Repository) GetByDayId(ctx context.Context, dayId string) []Media {
	var medias []Media
	repo.DataBase.DB.WithContext(ctx).
		Where("day_id = ?", dayId).
		Scan(&medias)
	return medias
}

// GetByDayId2Type получает медиа файл по ID дня и типу
func (repo *Repository) GetByDayId2Type(ctx context.Context, dayId string, mediaType string) (*Media, error) {
	var media Media
	result := repo.DataBase.DB.WithContext(ctx).
		Where("day_id = ?", dayId).
		Where("type = ?", mediaType).
		First(&media)
	if result.Error != nil {
		return nil, result.Error
	}
	return &media, nil
}

// Update обновляет запись в БД
func (repo *Repository) Update(ctx context.Context, media *Media) (*Media, error) {
	result := repo.DataBase.DB.WithContext(ctx).Clauses(clause.Returning{}).Updates(media)
	if result.Error != nil {
		return nil, result.Error
	}
	return media, nil
}
