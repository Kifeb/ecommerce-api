package main

import (
	productController "ecommerce_api/api/product/controller"
	productRepository "ecommerce_api/api/product/repository"
	productRoute "ecommerce_api/api/product/route"
	productService "ecommerce_api/api/product/service"
	"ecommerce_api/config"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func mainProduct() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db, _ := config.NewDB()
	validate := validator.New()
	productRepository := productRepository.NewProductRepository()
	productService := productService.NewProductService(productRepository, db, validate)
	productController := productController.NewProductController(productService)
	router := productRoute.NewProductRouter(productController)

	// router.HandleFunc("/v1", route.NewUserRouter(userController))

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	fmt.Println("Server Running")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
