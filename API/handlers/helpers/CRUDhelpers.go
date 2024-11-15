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

func Update(db *gorm.DB, w http.ResponseWriter, r *http.Request, info []byte, name string, table table) string {

	id := FindId(db, w, r, name, table)
	fmt.Println(id)
	switch v := table.(type) {
	case e.Product:
		Val := v
		err := json.Unmarshal(info, &Val)
		if err != nil {
			w.Write([]byte("Ошибка!\n"))
			fmt.Println(err)
			return "Всё плохо"
		}

		update := Val.NotNull()
		db.Table(name).Where("id = ?", id).Updates(update)
		s := ("Запись была обновлена")
		return s
	case e.ProductCategory:
		Val := v
		err := json.Unmarshal(info, &Val)
		if err != nil {
			w.Write([]byte("Ошибка!\n"))
			fmt.Println(err)
			return "Всё плохо"
		}

		update := Val.NotNull()
		db.Table(name).Where("id = ?", id).Updates(update)
		s := ("Запись была обновлена")
		return s
	default:
		return "Ничего не получилось"
	}
}

func Create(db *gorm.DB, w http.ResponseWriter, info []byte, name string, table table) {

	switch v := table.(type) {
	case e.Product:
		Val := v
		err := json.Unmarshal(info, &Val)
		if err != nil {
			w.Write([]byte("Ошибка!\n"))
			fmt.Println(err)
			return
		}
		db.Table(name).Create(&Val)
		w.Write([]byte("Создали запись!"))
	case e.ProductCategory:
		Val := v
		err := json.Unmarshal(info, &Val)
		fmt.Println(Val)
		if err != nil {
			w.Write([]byte("Ошибка!\n"))
			fmt.Println(err)
			return
		}
		db.Table(name).Create(&Val)
		w.Write([]byte("Создали запись!"))
	default:
		w.Write([]byte("Ошибка!\n"))
	}
}
