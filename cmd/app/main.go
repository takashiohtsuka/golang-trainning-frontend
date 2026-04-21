package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"

	"golang-trainning-frontend/pkg/config"
	"golang-trainning-frontend/pkg/infrastructure/datastore"
	"golang-trainning-frontend/pkg/infrastructure/router"
	"golang-trainning-frontend/pkg/infrastructure/validator"
	"golang-trainning-frontend/pkg/registry"
)

func main() {
	config.ReadConfig()

	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	db := datastore.NewDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	r := registry.NewRegistry(db)
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
