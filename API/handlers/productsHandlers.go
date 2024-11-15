package handlers

import (
	e "home/pavel/Go_tasks/API/entities"
	h_ "home/pavel/Go_tasks/API/handlers/helpers"
	"io"
	"net/http"
)

func ShowAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table []e.Product
	s := h_.ShowAll("product", table, db)
	w.Write([]byte(s))
}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.Product
	s := h_.Delete(db, r, "id = ?", "product", table)
	w.Write([]byte(s))
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.Product
	inf, _ := io.ReadAll(r.Body)

	s := h_.Update(db, w, r, inf, "product", table)
	w.Write([]byte(s))
}
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.Product
	inf, _ := io.ReadAll(r.Body)
	h_.Create(db, w, inf, "product", table)
}
