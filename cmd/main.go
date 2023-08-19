package main

import (
	"bia-challenge/internal/adapters/handler"
	"bia-challenge/internal/adapters/repository"
	"bia-challenge/internal/core/services"
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
	mysqlRepository := repository.NewMySQLRepository()
	biaService := services.NewBiaService(mysqlRepository)
	biaHandler := handler.NewHttpHandler(biaService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/consumption", biaHandler.GetEnergyConsumption)

	http.ListenAndServe(":3000", r)
}

func loadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
