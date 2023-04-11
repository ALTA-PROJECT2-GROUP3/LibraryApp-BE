package main

import (
	"libraryapp/app/config"
	"libraryapp/app/database"
	"libraryapp/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	db := database.InitDBMySql(*cfg)
	database.InitialMigration(db)

	e := echo.New()
	e.Use(middleware.CORS())
	router.InitRouter(db, e)
	e.Logger.Fatal(e.Start(":8080"))
}
