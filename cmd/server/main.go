package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/cmd/server/handler"
)

func main() {

	products := handler.NewProductHandler()
	if len(products) != 0 {
		router := chi.NewRouter()

		router.Get("/ping", handler.GetPing())

		router.Route("/products", func(r chi.Router) {
			r.Get("/", handler.GetAll(products))
			r.Get("/{id}", handler.GetByID(products))

		})

		http.ListenAndServe(":8080", router)

	}
}
