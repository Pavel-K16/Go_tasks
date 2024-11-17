package helpers

import (
	e "API/entities"
	d "API/env/envHelpers"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string

func MakeDsn() {
	dsn = d.Dsn()
}
func FindId(db *gorm.DB, w http.ResponseWriter, r *http.Request, name string, table e.Table) (uint, error) {
	var count int64
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Error: неверный данные в строке запроса")
		err := fmt.Errorf("неверный данные в строке запроса")
		return 0, err
	}
	db.Table(name).Where("id = ?", id).Count(&count) // .Error
	if count == 0 {
		w.Write([]byte("Запись с таким id не найдена"))
		err := fmt.Errorf("Запись в таблице" + name + "с  id: " + string(id) + "-- не найдена")
		return 0, err
	}
	return uint(id), nil
}
func Connection() (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Println("Не удалось подлючиться к БД", err)
		os.Exit(1)
	}
	conn, _ := db.DB()
	return db, conn
}
