package main

import (
	"fmt"

	"github.com/iYoarihtMacedo/GoWeb-GET-Practice/internal/application"
)

func main() {
	// env
	// app
	// - Config
	app := application.NewAppServer("localhost:8080")

	// - Run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return

	}
}
