package main

import (
	"aiCalendarApp/configs"
	"aiCalendarApp/internal/auth"
	"aiCalendarApp/internal/calendar/day"
	"aiCalendarApp/internal/calendar/media"
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
	dayRepository := day.NewRepository(db)
	mediaRepository := media.NewRepository(db)

	//Services
	dayService := day.NewService(*dayRepository)
	mediaService := media.NewService(*mediaRepository)

	// handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	day.NewHandler(router, day.HandlerDeps{
		Config:  conf,
		Service: dayService,
	})
	media.NewHandler(router, media.HandlerDeps{
		Config:  conf,
		Service: mediaService,
	})

	// server
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running on port 8081")
	server.ListenAndServe()
}
