package handlers

import (
	"fmt"
	e "home/pavel/Go_tasks/API/entities"
	h_ "home/pavel/Go_tasks/API/handlers/helpers"
	"io"
	"net/http"
)

func ShowAllCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table []e.ProductCategory
	s := h_.ShowAll("productcategory", table, db)
	w.Write([]byte(s))
}

func DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table_ e.ProductCategory
	h_.Delete(db, r, "categoryid = ?", "product", table_)

	var table e.ProductCategory
	h_.Delete(db, r, "id = ?", "productcategory", table)

	fmt.Fprintln(w, "запись успешно удалена")
}

func UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.ProductCategory
	inf, _ := io.ReadAll(r.Body)

	s := h_.Update(db, w, r, inf, "productcategory", table)
	w.Write([]byte(s))
}

func CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()

	var table e.ProductCategory
	inf, _ := io.ReadAll(r.Body)
	h_.Create(db, w, inf, "productcategory", table)
}
