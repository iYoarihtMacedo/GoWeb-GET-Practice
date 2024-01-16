package internal

import "errors"

var (
	ErrProductCodeAlreadyExists = errors.New("product already exists")
	ErrNotContent               = errors.New("no products existence")
	ErrIdNotExists              = errors.New("id does not exists")
)

type ProductRepository interface {
	Save(product *Product) (err error)
	GetAllProducts() (productList []Product, err error)
	GetById(id int) (product *Product, err error)
}
