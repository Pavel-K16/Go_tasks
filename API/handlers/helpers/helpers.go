package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	e "home/pavel/Go_tasks/API/entities"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type table interface {
	NotNull() map[string]interface{}
}

func Connection() (*gorm.DB, *sql.DB) {
	db, err := gorm.Open(postgres.Open(e.Dsn))
	if err != nil {
		panic(err.Error())
	}
	conn, _ := db.DB()
	return db, conn
}
func ShowAll(name string, table interface{}, db *gorm.DB) string {

	if err := db.Table(name).Find(&table).Error; err != nil {
		fmt.Println("Error fetching products:", err)
		return ""
	}
	info, _ := json.MarshalIndent(table, " ", "  ")
	s := ("Все записи таблицы " + name + ":\n" + string(info))
	return s
}

func Delete(db *gorm.DB, r *http.Request, id_s, name string, table interface{}) string {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := db.Table(name).Where(id_s, id).Delete(&table).Error; err != nil {
		log.Fatal("failed to delete product:", err)
	}
	s := "Запись успешно удалена"
	return s
}

func Update(db *gorm.DB, r *http.Request, name string, table table) string {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	result := db.Table(name).Find(&table, "id = ?", id)
	if result.RowsAffected == 0 {
		s := ("Запись с данным id не была найдена")
		return s
	}
	json.NewDecoder(r.Body).Decode(&table)
	update := table.NotNull()
	db.Table(name).Where("id = ?", id).Updates(update)
	s := ("Запись была обновлена")
	return s
}

func Create(db *gorm.DB, w http.ResponseWriter, r *http.Request, name string, table table) {

	json.NewDecoder(r.Body).Decode(&table) // Проверить пустое тело запроса или нет
	db.Table(name).Create(&table)
	w.Write([]byte("Создали запись!"))
}
