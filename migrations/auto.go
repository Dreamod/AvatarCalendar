package main

import (
	"aiCalendarApp/configs"
	"aiCalendarApp/internal/calendar/day"
	"aiCalendarApp/internal/calendar/media"
	"aiCalendarApp/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	db.AutoMigrate(&day.Day{})
	db.AutoMigrate(&media.Media{})
}
