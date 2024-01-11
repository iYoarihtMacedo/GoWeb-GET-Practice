package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal/product"
)

func main() {

	products := []product.Product{}

	data, err := os.ReadFile("./products.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := json.Unmarshal(data, &products); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(products)

}
