package db

import (
	"aiCalendarApp/configs"
)
import "gorm.io/gorm"
import "gorm.io/driver/postgres"

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
