package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" ///!!!!!!!!!!
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type productCategory struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:VARCHAR(25);not null"`
	Description string `gorm:"type:text;column:description"`
}

func (item *productCategory) NotNull() map[string]interface{} {
	toupdate := make(map[string]interface{})

	if item.Name != "" {
		toupdate["name"] = item.Name
	}
	if item.Description != "" {
		toupdate["description"] = item.Description
	}
	return toupdate
}

type product struct {
	Id          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"type:VARCHAR(25);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:Decimal(15,3)"`
	CategoryId  int     `gorm:"foreignKey; references: productCategory; Id;column:categoryid"`
}

// curl -X POST -H "Content-Type: application/json" -d '{"Name":"John Doe","Description":"john@example.com", "Price":23443,"categoryid":1}' http://localhost:8080/item
//curl -X PUT -H "Content-Type: application/json" -d '{"Name":"Iphone 16", "description":"Очень хороший телефон", "Price": 96500,"categoryid":1}' http://localhost:8080/item/2

func (item *product) NotNull() map[string]interface{} {
	toupdate := make(map[string]interface{})

	if item.Name != "" {
		toupdate["name"] = item.Name
	}
	if item.Description != "" {
		toupdate["description"] = item.Description
	}
	if item.Price != 0 {
		toupdate["price"] = item.Price
	}
	if item.CategoryId != 0 {
		toupdate["categoryid"] = item.CategoryId
	}
	return toupdate
}

const (
	dbHost = "localhost"
	dbUser = "pavel"
	dbPass = "55544"
	dbName = "shop"
	dbPort = "5432"
	dbSsl  = "disable"
	dsn    = "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSsl
	serv   = ":8080"
)

// http://localhost:8080/item/1
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/item/{id}", getInfoHandler).Methods("GET")
	r.HandleFunc("/item", createProductHandler).Methods("POST")
	r.HandleFunc("/category", createCategoryHandler).Methods("POST")
	r.HandleFunc("/item/{id}", updateProductHandler).Methods("PUT")
	r.HandleFunc("/category/{id}", updateCategoryHandler).Methods("PUT")
	r.HandleFunc("/item/{id}", deleteProductHandler).Methods("DELETE")
	r.HandleFunc("/category/{id}", deleteCategoryHandler).Methods("DELETE")

	err := http.ListenAndServe(serv, r)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера", err)
	}
}
func deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	conn, _ := db.DB()
	defer conn.Close()

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err = db.Table("productCatagory").Where("id = ?", id).Delete(&productCategory{}).Error; err != nil {
		log.Fatal("failed to delete product category:", err)
	}

	if err = db.Table("product").Where("categoryid = ?", id).Delete(&product{}).Error; err != nil {
		log.Fatal("failed to delete product:", err)
	}

	fmt.Fprintln(w, "запись успешно удалена")

}
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	conn, _ := db.DB()
	defer conn.Close()

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err = db.Table("product").Where("id = ?", id).Delete(&product{}).Error; err != nil {
		log.Fatal("failed to delete product:", err)
	}
	fmt.Fprintln(w, "запись успешно удалена")
}
func updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	conn, _ := db.DB()
	defer conn.Close()
	var item productCategory
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	json.NewDecoder(r.Body).Decode(&item)
	update := item.NotNull()
	db.Model(&product{}).Where("id = ?", id).Updates(update)

	w.Write([]byte("Всё прошло успешно, изменения были приняты"))
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	conn, _ := db.DB()
	defer conn.Close()
	var item product
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	json.NewDecoder(r.Body).Decode(&item)
	update := item.NotNull()
	db.Model(&product{}).Table("product").Where("id = ?", id).Updates(update)

	w.Write([]byte("Всё прошло успешно, изменения были приняты"))
}

func createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	conn, _ := db.DB()
	defer conn.Close()

	var category productCategory
	json.NewDecoder(r.Body).Decode(&category)
	db.Table("productcategory").Create(&category)
	w.Write([]byte("ergergerg"))

}
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err.Error())
	}
	conn, _ := db.DB()
	defer conn.Close()

	var item product
	json.NewDecoder(r.Body).Decode(&item)

	db.Table("product").Create(&item)
	w.Write([]byte("ergergerg"))
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Не удалось подключиться к БД", err)
	}
	conn, _ := db.DB()
	defer conn.Close()
	var item product
	var catitem productCategory
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	db.Table("product").Find(&item, "id = ?", id)
	fmt.Println(item)
	db.Table("productcategory").Find(&catitem, "id = ?", item.CategoryId) //!!!

	fmt.Println(catitem)
	w.Write([]byte("Наименование товара:\n"))
	info, _ := json.MarshalIndent(item, "", "  ")
	w.Write(info)
	info_, err := json.MarshalIndent(catitem, "", "  ")
	if err != nil {
		fmt.Println("НЕ маршалит категорию")
	}

	w.Write([]byte("\nКатегория товара:\n"))
	w.Write(info_)
}
