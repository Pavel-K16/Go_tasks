package handlers

import (
	e "API/entities"
	h_ "API/handlers/helpers"
	"encoding/json"
	"fmt"
	"log"
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
	if id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error: неверный данные в строке запроса")
		return
	}
	result := db.Table("product").Find(&ptable, "id = ?", id)
	if result.RowsAffected == 0 {
		w.Write([]byte("Запись с данным id не была найдена"))
		return
	}

	info, _ := json.MarshalIndent(ptable, "", "  ")

	db.Table("productcategory").Find(&ctable, "id = ?", ptable.CategoryId) //!!!
	info_, err := json.MarshalIndent(ctable, "", "  ")
	if err != nil {
		log.Println("Ошибка при декодировании данных")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка при декодировании данных"))
		return
	}

	w.Write([]byte("Наименование товара:\n"))
	w.Write(info)
	w.Write([]byte("\nКатегория товара:\n"))
	w.Write(info_)
}
