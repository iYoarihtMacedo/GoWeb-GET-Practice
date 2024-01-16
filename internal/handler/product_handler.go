package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal"
	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/platform/web/response"
)

func NewDefaultProducts(sv internal.ProductService) *DefaultProducts {
	return &DefaultProducts{
		sv: sv,
	}
}

type DefaultProducts struct {
	sv internal.ProductService
}

func (p *DefaultProducts) GetPing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.Text(w, http.StatusOK, "pong")
	}
}

func (p *DefaultProducts) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := (*p).sv.GetAllProducts()
		if err != nil {
			response.JSON(w, http.StatusNoContent, map[string]any{
				"message": "Data not found",
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Data found",
			"data":    data,
		})

	}
}

func (p *DefaultProducts) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "Something went wrong",
			})
		}

		prod, err := (*p).sv.GetById(id)
		if err != nil {
			response.JSON(w, http.StatusNoContent, map[string]any{
				"message": "Data not found",
			})
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Product found",
			"data":    prod,
		})

	}
}

func (p *DefaultProducts) AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
