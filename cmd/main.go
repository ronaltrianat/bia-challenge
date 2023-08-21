package main

import (
	mysqlhelper "bia-challenge/cmd/helpers/mysql"
	"bia-challenge/internal/adapters/handler"
	"bia-challenge/internal/adapters/repository"
	addressesservice "bia-challenge/internal/core/services/addresses"
	consumptionservice "bia-challenge/internal/core/services/consumption"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main() {
	loadConfiguration()
	initRoutes()
}

func initRoutes() {
	mysqlDB := mysqlhelper.GetMySQLConnection()
	mysqlRepository := repository.NewMySQLRepository(mysqlDB)
	addressesService := addressesservice.NewAddressesService()
	biaService := consumptionservice.NewBiaService(mysqlRepository, addressesService)
	biaHandler := handler.NewHttpHandler(biaService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/consumption", biaHandler.GetEnergyConsumption)

	log.Println("starting app")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func loadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
