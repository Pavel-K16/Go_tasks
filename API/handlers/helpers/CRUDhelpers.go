package helpers

import (
	"encoding/json"
	"fmt"
	e "home/pavel/Go_tasks/API/entities"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func ShowAll(name string, table interface{}, db *gorm.DB) string {

	if err := db.Table(name).Find(&table).Error; err != nil {
		fmt.Println("Error fetching products:", err)
		return ""
	}
	info, err := json.MarshalIndent(table, " ", "  ")
	if err != nil {
		log.Println("Ошибка при диссерелизации данных")
		return ""
	}
	s := ("Все записи таблицы " + name + ":\n" + string(info))
	return s
}

func Delete(db *gorm.DB, r *http.Request, id_s, name string, table interface{}) string {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := db.Table(name).Where(id_s, id).Delete(&table).Error; err != nil {
		log.Println("failed to delete product:", err)
		return ""
	}
	s := "Запись успешно удалена"
	return s
}

func Update(db *gorm.DB, w http.ResponseWriter, r *http.Request, info []byte, name string, table e.Table) string {

	id, err := FindId(db, w, r, name, table)
	if err != nil {
		return ""
	}
	switch v := table.(type) {
	case e.Product:
		Val := v
		err := json.Unmarshal(info, &Val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Переданы некорректные данные\n"))
			return ""
		}

		update := Val.NotNull()
		db.Table(name).Where("id = ?", id).Updates(update)
		s := ("Запись c id: " + string(id) + " была обновлена")
		return s
	case e.ProductCategory:
		Val := v
		err := json.Unmarshal(info, &Val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Переданы некорректные данные\n"))
			return ""
		}

		update := Val.NotNull()
		db.Table(name).Where("id = ?", id).Updates(update)
		s := ("Запись c id: " + string(id) + " была обновлена")
		return s
	default:
		return ""
	}
}

func Create(db *gorm.DB, w http.ResponseWriter, info []byte, name string, table e.Table) {

	switch v := table.(type) {
	case e.Product:
		Val := v
		err := json.Unmarshal(info, &Val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Переданы некорректные данные\n"))
			log.Println("Переданы некорректные данные")
			return
		}
		if Val.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Попытка создать запись без имени"))
			return
		}
		db.Table(name).Create(&Val)
		w.Write([]byte("Создали запись!"))
	case e.ProductCategory:
		Val := v
		err := json.Unmarshal(info, &Val)
		fmt.Println(Val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Переданы некорректные данные\n"))
			log.Println("Переданы некорректные данные")
			return
		}
		if Val.Name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Попытка создать запись без имени"))
			return
		}
		db.Table(name).Create(&Val)
		w.Write([]byte("Создали запись!"))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Ошибка сервера!\n"))
		log.Println("Ошибка сервера")
	}
}
