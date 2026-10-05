package day

import (
	"aiCalendarApp/pkg/db"
	"fmt"

	"gorm.io/gorm/clause"
)

type DayRepository struct {
	DataBase *db.Db
}

func NewDayRepository(dataBase *db.Db) *DayRepository {
	return &DayRepository{
		DataBase: dataBase,
	}
}

func (repo *DayRepository) Create(day *Day) (*Day, error) {
	result := repo.DataBase.DB.Create(day)
	if result.Error != nil {
		return nil, result.Error
	}
	return day, nil
}

func (repo *DayRepository) GetById(id string) (*Day, error) {
	var day Day
	result := repo.DataBase.DB.First(&day, "id = ?", id)
	fmt.Println(result)
	if result.Error != nil {
		return nil, result.Error
	}
	return &day, nil
}

func (repo *DayRepository) Update(day *Day) (*Day, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(day)
	if result.Error != nil {
		return nil, result.Error
	}
	return day, nil
}
