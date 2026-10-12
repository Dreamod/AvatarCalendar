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

//TODO запуск миграций на старте приложения, изучить вопрос, мб стоит делать миграции только при компиляции или еще как
