package service

import (
	"fmt"

	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal"
)

// NewProductDefault creates a new instance of a product service
func NewProductDefault(rp internal.ProductRepository) *ProductDefault {
	return &ProductDefault{
		rp: rp,
	}
}

// ProductDefault is a struct that represents the default implementation of a movie service
type ProductDefault struct {
	// rp is a product repository
	rp internal.ProductRepository
	// external services
	// ... (weather api, etc)
}

func (p *ProductDefault) GetAllProducts() (productList []internal.Product, err error) {

	// fmt.Println("Service:", (*p).rp)
	products, err := (*p).rp.GetAllProducts()
	if err != nil {
		return products, internal.ErrNoProducts
	}

	productList = products

	return
}

func (p *ProductDefault) GetById(id int) (product *internal.Product, err error) {

	prod, err := (*p).rp.GetById(id)
	if err != nil {
		return &internal.Product{}, internal.ErrIdDoesntExists
	}

	product = prod

	return

}

// Save saves a product validating its data
func (p *ProductDefault) Save(product *internal.Product) (err error) {
	// External services
	// ...

	// business logic
	// - validation of required fields
	if (*product).CodeValue == "" {
		return fmt.Errorf("%w: code value", internal.ErrFieldRequired)
	}

	if (*product).Expiration == "" {
		return fmt.Errorf("%w: expiration", internal.ErrFieldRequired)
	}

	return
}
