package main

import (
	"ecommerce_api/api/users/controller"
	"ecommerce_api/api/users/repository"
	"ecommerce_api/api/users/route"
	"ecommerce_api/api/users/service"
	"ecommerce_api/config"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	db, _ := config.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := route.NewUserRouter(userController)

	// productRepository := productRepository.NewProductRepository()
	// productService := productService.NewProductService(productRepository, db, validate)
	// productController := productController.NewProductController(productService)
	// router = productRoute.NewProductRouter(productController)

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
