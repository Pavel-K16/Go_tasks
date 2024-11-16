package handlers

import (
	e "home/pavel/Go_tasks/API/entities"
	h_ "home/pavel/Go_tasks/API/handlers/helpers"
	"io"
	"log"
	"net/http"
)

func ShowAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table []e.Product
	s := h_.ShowAll("product", table, db)
	if s == "" {
		w.WriteHeader(http.StatusInternalServerError)
		s = "Ошибка при обработке данных из БД"
	}
	w.Write([]byte(s))
}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.Product
	s := h_.Delete(db, r, "id = ?", "product", table)
	if s == "" {
		s = "Не удалось удалить запись"
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write([]byte(s))
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.Product
	inf, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Не удалось прочитать тело запроса", err)
	}
	s := h_.Update(db, w, r, inf, "product", table)
	w.Write([]byte(s))
}
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.Product
	inf, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Не удалось прочитать тело запроса", err)
	}
	h_.Create(db, w, inf, "product", table)
}
