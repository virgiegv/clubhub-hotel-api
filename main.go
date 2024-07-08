package main

import (
	_ "clubhub-hotel-api/docs"
	db2 "clubhub-hotel-api/repository/models"
	"github.com/labstack/echo/v4"
)

// @title Clubhub Hotel Franchise API
// @version 1.0
// @description This API manages the CRU operations of Clubhub's hotel franchises

// @contact.name API Support
// @contact.url https://github.com/virgiegv/
// @contact.email vgil_22@hotmail.com

func main() {
	e := echo.New()

	db2.Init()

	router := Router{server: e}
	router.Init()

	e.Logger.Fatal(e.Start(":8080"))

}
