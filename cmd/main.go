package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/nao50/ddd-api/config"
	"github.com/nao50/ddd-api/infrastructure"
	"github.com/nao50/ddd-api/interface/handler"
	"github.com/nao50/ddd-api/usecase"
)

func main() {
	petInfrastructure := infrastructure.NewPetRepository(config.NewDB())
	petUsecase := usecase.NewPetUsecase(petInfrastructure)
	petHandler := handler.NewPetHandler(petUsecase)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/pets", petHandler.ListPets)
	r.Get("/pets/{id}", petHandler.GetPet)
	r.Post("/pets", petHandler.CreatePet)
	r.Put("/pets/{id}", petHandler.UpdatePet)
	r.Delete("/pets/{id}", petHandler.DeletePet)

	http.ListenAndServe(":5050", r)
}
