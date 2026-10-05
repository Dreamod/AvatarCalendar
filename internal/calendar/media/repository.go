package media

import (
	"aiCalendarApp/pkg/db"
	"fmt"

	"gorm.io/gorm/clause"
)

type MediaRepository struct {
	DataBase *db.Db
}

func NewMediaRepository(dataBase *db.Db) *MediaRepository {
	return &MediaRepository{
		DataBase: dataBase,
	}
}

func (repo *MediaRepository) Create(media *Media) (*Media, error) {
	result := repo.DataBase.DB.Create(media)
	if result.Error != nil {
		return nil, result.Error
	}
	return media, nil
}

func (repo *MediaRepository) GetById(id string) (*Media, error) {
	var media Media
	result := repo.DataBase.DB.First(&media, "id = ?", id)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &media, nil
}

func (repo *MediaRepository) Update(media *Media) (*Media, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(media)
	if result.Error != nil {
		return nil, result.Error
	}
	return media, nil
}
