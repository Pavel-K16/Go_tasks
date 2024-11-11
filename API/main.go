package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux" ///!!!!!!!!!!
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*type ProductCategory struct {
	Id   uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(10)"`
    Descritpion string `gorm:"type:text"`
}*/

type phones struct {
	Id          int     `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"type:varchar(10);not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:Decimal(10,3)"`
	//CategoryId uint `gorm:"foreignKey; references: ProductCategory; Id"`
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

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/phones/{id}", getInfoHandler).Methods("GET")

	err := http.ListenAndServe(serv, r)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера", err)
	}
	/*
	   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	   	if err != nil {
	   		log.Fatal("Ошибка при подключении к БД Shop")
	   	}

	   db.AutoMigrate(&jopa{})

	   data := jopa{Name: "Smartphone", Description: "The latest model smartphone", Price: 699.99}
	   db.Create(&data)
	   var products []jopa
	   db.Find(&products)

	   	for _, product := range products {
	   		fmt.Println(product.Id, product.Name, product.Description, product.Price)
	   	}

	   /*response, err := http.Get("https://stepik.org/lesson/1101589/step/7?auth=login&unit=1112589")

	   	if err != nil {
	   		fmt.Print("JOPA WITH GET REQUEST")
	   	}

	   //info, _ := io.ReadAll(response.Body)

	   fmt.Println(response.Status)
	*/
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(postgres.Open(dsn)) //???
	if err != nil {
		log.Fatal("Не удалось подключиться к БД", err)
	}
	conn, _ := db.DB()
	defer conn.Close()
	var phone phones
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	db.Find(&phone, "id = ?", id)
	info, _ := json.MarshalIndent(phone, "", "  ")
	w.Write(info)
}
