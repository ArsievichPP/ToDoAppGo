package main

import (
	"ToDoAppGo/app/api"
	"ToDoAppGo/app/config"
	"ToDoAppGo/app/internal/db"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	// Read Config
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize Echo
	e := echo.New()

	// Connect to DB
	database, err := db.GetInstance(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Routes
	api.InitRoutes(e, database)

	//Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
