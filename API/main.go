package main

import (
	"log"
	"net/http"

	e "home/pavel/Go_tasks/API/entities"
	h "home/pavel/Go_tasks/API/handlers"

	"github.com/gorilla/mux" ///!!!!!!!!!!
)

// curl -X POST -H "Content-Type: application/json" -d '{"Name":"John Doe","Description":"john@example.com", "Price":23443,"categoryid":1}' http://localhost:8080/item
//curl -X PUT -H "Content-Type: application/json" -d '{"Name":"Iphone 16", "description":"Очень хороший телефон", "Price": 96500,"categoryid":1}' http://localhost:8080/item/2

// http://localhost:8080/item/1
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/item/{id}", h.GetInfoHandler).Methods("GET")
	r.HandleFunc("/item", h.CreateProductHandler).Methods("POST")
	r.HandleFunc("/category", h.CreateCategoryHandler).Methods("POST")
	r.HandleFunc("/item/{id}", h.UpdateProductHandler).Methods("PUT")
	r.HandleFunc("/category/{id}", h.UpdateCategoryHandler).Methods("PUT")
	r.HandleFunc("/item/{id}", h.DeleteProductHandler).Methods("DELETE")
	r.HandleFunc("/category/{id}", h.DeleteCategoryHandler).Methods("DELETE")
	r.HandleFunc("/item", h.ShowAllProductsHandler).Methods("GET")
	r.HandleFunc("/category", h.ShowAllCategoryHandler).Methods("GET")
	err := http.ListenAndServe(e.Serv, r)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера", err)
	}
}
