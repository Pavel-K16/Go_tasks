package main

import (
	h "API/handlers"
	db "API/handlers/helpers"

	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// curl -X POST -H "Content-Type: application/json" -d '{"Name":"John Doe","Description":"john@example.com", "Price":23443,"categoryid":1}' http://localhost:8080/item
// curl -X PUT -H "Content-Type: application/json" -d '{"Name":"Iphone 16", "description":"Очень хороший телефон", "Price": 96500,"categoryid":1}' http://localhost:8080/item/2
// curl -X DELETE http://localhost:8080/item/20
// http://localhost:8080/item/1

// curl -X POST -H "Content-Type: application/json" -d '{"Name":"John Doe","Description":"john@example.com"}' http://localhost:8080/category
// curl -X PUT -H "Content-Type: application/json" -d '{"Name":"Дорогие телефоны", "description":"В этой категории телефоны дороже 100к"}' http://localhost:8080/category/2
// curl -X DELETE http://localhost:8080/category/2

func main() {

	err := godotenv.Load("env/dsn.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.MakeDsn()

	r := mux.NewRouter()

	r.HandleFunc("/item/{id}", h.GetInfoHandler).Methods("GET")
	r.HandleFunc("/item", h.ShowAllProductsHandler).Methods("GET")
	r.HandleFunc("/category", h.ShowAllCategoryHandler).Methods("GET")

	r.HandleFunc("/item", h.CreateProductHandler).Methods("POST")
	r.HandleFunc("/category", h.CreateCategoryHandler).Methods("POST")

	r.HandleFunc("/item/{id}", h.UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/category/{id}", h.UpdateCategoryHandler).Methods("PUT")

	r.HandleFunc("/item/{id}", h.DeleteProductHandler).Methods("DELETE")
	r.HandleFunc("/category/{id}", h.DeleteCategoryHandler).Methods("DELETE")
	serv := os.Getenv("SERV")
	err = http.ListenAndServe(serv, r)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера", err)
	}
}
