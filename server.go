package main

import (
	"github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	pDelivery "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/delivery"
	uDelivery "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/delivery"

	pRepo "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/repository"
	uRepo "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/repository"

	pServices "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/posting/services"
	uServices "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/features/user/services"

	database "github.com/ALTA-Group-Project-Social-Media-Apps/Social-Media-Apps/utils"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := database.InitDB(cfg)
	database.MigrateDB(db)
	uRepo := uRepo.New(db)
	pRepo := pRepo.New(db)

	uService := uServices.New(uRepo)
	pService := pServices.New(pRepo)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	uDelivery.New(e, uService)
	pDelivery.New(e, pService)

	e.Logger.Fatal(e.Start(":8000"))
}
