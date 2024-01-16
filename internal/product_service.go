package internal

import "errors"

var (
	ErrFieldRequired        = errors.New("field required")
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrNoProducts           = errors.New("no products existence")
	ErrIdDoesntExists       = errors.New("product id does not exists")
)

// ProductService is an interface that represents product service
// - business logic
// - validation
// - external services (e.g. apis, db, etc.)
type ProductService interface {
	Save(product *Product) (err error)
	GetAllProducts() (productList []Product, err error)
	GetById(id int) (product *Product, err error)
}
