package handlers

import (
	e "API/entities"
	h_ "API/handlers/helpers"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ShowAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table []e.ProductCategory
	s := h_.ShowAll("productcategory", table, db)
	if s == "" {
		w.WriteHeader(http.StatusInternalServerError)
		s = "Ошибка при обработке данных из БД"
	}
	w.Write([]byte(s))
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table_ e.ProductCategory
	s := h_.Delete(db, r, "categoryid = ?", "product", table_)
	if s == "" {
		s = "Не удалось удалить запись"
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var table e.ProductCategory
	s = h_.Delete(db, r, "id = ?", "productcategory", table)
	if s == "" {
		s = "Не удалось удалить запись"
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "запись успешно удалена")
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.ProductCategory
	inf, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Не удалось прочитать тело запроса", err)
	}
	s := h_.Update(db, w, r, inf, "productcategory", table)
	w.Write([]byte(s))
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.ProductCategory
	inf, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Не удалось прочитать тело запроса", err)
	}
	h_.Create(db, w, inf, "productcategory", table)
}
