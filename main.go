package main

import (
	productController "ecommerce_api/api/product/controller"
	productRepository "ecommerce_api/api/product/repository"
	productService "ecommerce_api/api/product/service"
	"ecommerce_api/api/users/controller"
	"ecommerce_api/api/users/repository"
	"ecommerce_api/api/users/service"
	"ecommerce_api/config"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load file .env untuk didistribusikan
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	// deklasrasi mux router
	mux := http.DefaultServeMux

	// insialisasi koneksi databases
	db, _ := config.NewDB()

	// instansiasi validate untuk setiap request dari user
	validate := validator.New()

	// Repository Pattern untuk User API
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	// Repository Pattern untuk Product API
	productRepository := productRepository.NewProductRepository()
	productService := productService.NewProductService(productRepository, db, validate)
	productController := productController.NewProductController(productService)

	// Registrasi setiap user dan panggil masing controller yang bersagnkutan
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userId")
		productId := r.URL.Query().Get("productId")
		uid, _ := strconv.Atoi(userId)
		pid, _ := strconv.Atoi(productId)

		switch r.Method {
		case "POST":
			userController.Create(w, r)
		case "GET":
			if uid > 0 && pid > 0 {
				userController.GetProducyById(w, r, uid, pid)
			} else if uid > 0 && pid == 0 {
				userController.GetProductByUser(w, r, uid)
			} else {
				userController.FindAll(w, r)
			}
		case "PUT":
			userController.UpdateProductBySeller(w, r, pid, uid)
		default:
			http.Error(w, "Kosong", http.StatusBadRequest)
		}
	})

	mux.HandleFunc("/api/update", func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userId")
		productId := r.URL.Query().Get("productId")
		uid, _ := strconv.Atoi(userId)
		pid, _ := strconv.Atoi(productId)

		fmt.Println(uid)
		userController.UpdateProductBySeller(w, r, uid, pid)
	})

	mux.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		uid := r.URL.Query().Get("userId")
		userId, _ := strconv.Atoi(uid)
		switch r.Method {
		case "POST":
			productController.Create(w, r, userId)
		case "GET":
			productController.FindAll(w, r)
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	mux.HandleFunc("/api/users/purchase", func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("userId")
		productId := r.URL.Query().Get("productId")
		uid, _ := strconv.Atoi(userId)
		pid, _ := strconv.Atoi(productId)

		userController.Purchase(w, r, uid, pid)
	})

	// Running Server
	var handler http.Handler = mux
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: handler,
	}

	fmt.Println("Server Running")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
