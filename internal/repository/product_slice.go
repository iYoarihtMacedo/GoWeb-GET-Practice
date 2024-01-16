package repository

import (
	"encoding/json"
	"fmt"
	"os"

	internal "github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal"
)

func NewProductSlice() *ProductSlice {
	data, err := os.ReadFile("./docs/db/products.json")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	productStorage := []internal.Product{}
	idCounter := 0

	if err := json.Unmarshal(data, &productStorage); err != nil {
		fmt.Println(err)
		return nil
	}

	for i := 0; i < len(productStorage); i++ {
		idCounter++
	}

	ps := ProductSlice{
		db:     productStorage,
		lastId: idCounter,
	}

	return &ps
}

// ProducSlice represents the product storage
type ProductSlice struct {
	// db is the slice that contains all the product storage
	db []internal.Product
	// lastId gets the last id assigned to the last product inserted in the product slice
	lastId int
}

func (p *ProductSlice) GetAllProducts() (productList []internal.Product, err error) {

	// fmt.Println("Repository:", (*p).db)
	if len((*p).db) == 0 {
		return []internal.Product{}, internal.ErrNotContent
	}

	productList = (*p).db

	return
}

func (p *ProductSlice) GetById(id int) (product *internal.Product, err error) {

	if id > p.lastId {
		return &internal.Product{}, internal.ErrIdNotExists
	}

	for _, prod := range (*p).db {
		if prod.ID == id {
			product = &prod
			break
		}
	}

	return
}

func (p *ProductSlice) Save(product *internal.Product) (err error) {

	// Validate product (consistency)
	// - codevalue: unique
	for _, v := range (*p).db {
		if v.CodeValue == product.CodeValue {
			return internal.ErrProductCodeAlreadyExists
		}
	}

	// autoincrement to last id
	(*p).lastId++

	// assignment last id to new product
	(*product).ID = (*p).lastId

	// storing new product
	(*p).db[(*product).ID] = *product

	return

}
