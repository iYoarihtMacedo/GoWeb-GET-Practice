package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	product "github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal/domain"
)

type Res struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewProductHandler() []product.Product {

	products := []product.Product{}

	data, err := os.ReadFile("./cmd/server/handler/products.json")
	if err != nil {
		fmt.Println(err)
		return products
	}

	if err := json.Unmarshal(data, &products); err != nil {
		fmt.Println(err)
		return products
	}

	return products
}

func GetPing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

func GetAll(p []product.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if len(p) == 0 {
			code := http.StatusNotFound
			body := Res{
				Message: "Data not Found",
				Data:    nil,
			}

			fmt.Printf("Log: %d - Data not Found", code)
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)

			return
		}

		code := http.StatusOK
		body := Res{
			Message: "Data Found",
			Data:    p,
		}

		fmt.Printf("Log: %d - Data Found", code)
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}
}

func GetByID(p []product.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			fmt.Println(err)
			return
		}

		var data product.Product

		pSize := len(p)

		if pSize == 0 {
			code := http.StatusNotFound
			body := Res{
				Message: "Data not Found",
				Data:    nil,
			}

			fmt.Printf("Log: %d - Data not Found", code)
			w.WriteHeader(code)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(body)

			return
		}

		for i := 0; i < pSize; i++ {
			if p[i].ID == id {
				data = p[i]
				break
			}
		}

		code := http.StatusOK
		body := Res{
			Message: "Data Found",
			Data:    data,
		}

		fmt.Printf("Log: %d - Data Found", code)
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)

	}
}
