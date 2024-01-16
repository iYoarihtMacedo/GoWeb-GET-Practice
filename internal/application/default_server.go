package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal/handler"
	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal/repository"
	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal/service"
)

type AppServer struct {
	address string
}

func NewAppServer(address string) *AppServer {

	defaultAddress := ":8080"
	if address != "" {
		defaultAddress = address

	}

	return &AppServer{
		address: defaultAddress,
	}

}

func (a *AppServer) Run() error {
	// load dependencies
	// 1) repository
	db := repository.NewProductSlice()

	// 2) service
	sv := service.NewProductDefault(db)

	// 3) handler
	hd := handler.NewDefaultProducts(sv)

	// 4) Router
	rt := chi.NewRouter()

	//	4.1) Endpoints
	rt.Get("/ping", hd.GetPing())
	rt.Route("/products", func(r chi.Router) {
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetById())
		r.Post("/", hd.AddProduct())
	})

	// 5) Run server
	return http.ListenAndServe(a.address, rt)
}
