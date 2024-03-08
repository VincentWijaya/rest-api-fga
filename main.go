package main

import (
	"rest-api/database"
	"rest-api/handler"
	productHandler "rest-api/handler/product"
	"rest-api/repo/product"
	productSvc "rest-api/service/product"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		panic(err)
	}

	// Setup Repo
	productRepo := product.NewRepository(db)

	// Setup Service
	productService := productSvc.NewService(productRepo)

	// Setup Handler
	productHandler := productHandler.NewHandler(productService)

	handler.NewHttpServer(productHandler)
}
