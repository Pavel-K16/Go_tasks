package handlers

import (
	"encoding/json"
	"fmt"
	e "home/pavel/Go_tasks/API/entities"
	h_ "home/pavel/Go_tasks/API/handlers/helpers"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetInfoHandler(w http.ResponseWriter, r *http.Request) {
	db, conn := h_.Connection()
	defer conn.Close()
	var ptable e.Product
	var ctable e.ProductCategory

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	result := db.Table("product").Find(&ptable, "id = ?", id)
	if result.RowsAffected == 0 {
		w.Write([]byte("Запись с данным id не была найдена"))
		return
	}

	info, _ := json.MarshalIndent(ptable, "", "  ")

	db.Table("productcategory").Find(&ctable, "id = ?", ptable.CategoryId) //!!!
	info_, err := json.MarshalIndent(ctable, "", "  ")
	if err != nil {
		fmt.Println("НЕ маршалит категорию")
	}

	w.Write([]byte("Наименование товара:\n"))
	w.Write(info)
	w.Write([]byte("\nКатегория товара:\n"))
	w.Write(info_)
}
