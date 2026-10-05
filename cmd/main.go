package main

import (
	"aiCalendarApp/configs"
	"aiCalendarApp/internal/auth"
	"aiCalendarApp/internal/calendar/day"
	"aiCalendarApp/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	// config
	conf := configs.LoadConfig()

	// db
	db := db.NewDb(conf)

	// router
	router := http.NewServeMux()

	//Repositories
	dayRepository := day.NewDayRepository(db)

	// handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	day.NewDayHandler(router, day.DayHandlerDeps{
		Config:        conf,
		DayRepository: dayRepository,
	})

	// server
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running on port 8081")
	server.ListenAndServe()
}
