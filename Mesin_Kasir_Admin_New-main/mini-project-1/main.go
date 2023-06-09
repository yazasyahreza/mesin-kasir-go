package main

import (
	"fmt"
	"mini-project/middleware"
	"mini-project/modules/admin"
	"mini-project/modules/products"
	"mini-project/modules/transactions"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/last_project?parseTime=true"))
	if err != nil {
		panic("Failed to connect database")
	}

	productRepo := products.Repository{DB: db}
	productUsecase := products.Usecase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUsecase}

	transactionRepo := transactions.Repository{DB: db}
	transactionUsecase := transactions.Usecase{TransacationRepo: transactionRepo, ProductRepo: productRepo}
	transactionHandler := transactions.Handler{Usecase: transactionUsecase}

	adminRepo := admin.Repository{DB: db}
	adminUsecase := admin.Usecase{Repo: adminRepo}
	admiHandler := admin.Handler{Usecase: adminUsecase}

	router := mux.NewRouter()

	router.HandleFunc("/admin/login", admiHandler.Login).Methods("POST")

	router.HandleFunc("/products", middleware.MiddlewareJWTAuthorization(productHandler.GetAllProducts)).Methods("GET")
	router.HandleFunc("/products/{id}", middleware.MiddlewareJWTAuthorization(productHandler.GetProductById)).Methods("GET")
	router.HandleFunc("/products", middleware.MiddlewareJWTAuthorization(productHandler.AddProduct)).Methods("POST")
	router.HandleFunc("/products/{id}", middleware.MiddlewareJWTAuthorization(productHandler.EditProduct)).Methods("PUT")
	router.HandleFunc("/products/{id}/status", middleware.MiddlewareJWTAuthorization(productHandler.SoftDelete)).Methods("PATCH")

	router.HandleFunc("/transactions", middleware.MiddlewareJWTAuthorization(transactionHandler.GetAll)).Methods("GET")
	router.HandleFunc("/transactions/{id}", middleware.MiddlewareJWTAuthorization(transactionHandler.GetById)).Methods("GET")
	router.HandleFunc("/transactions", middleware.MiddlewareJWTAuthorization(transactionHandler.Create)).Methods("POST")

	PORT := ":8080"
	fmt.Println("Starting server at localhost", PORT)
	http.ListenAndServe(PORT, router)
}
