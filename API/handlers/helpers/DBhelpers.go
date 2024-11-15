package helpers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func FindId(db *gorm.DB, w http.ResponseWriter, r *http.Request, name string, table table) uint {
	var count int64
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	db.Table(name).Where("id = ?", id).Count(&count) // .Error
	if count == 0 {
		w.Write([]byte("Запись с таким id не найдена"))
		return 0
	}
	return uint(id)
}
