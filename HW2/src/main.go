package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Number struct {
	Id     int `gorm:"primaryKey"`
	Number int `gorm:"unique" json:"number"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var DB *gorm.DB

func main() {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=sem password=sem dbname=sem port=5432 TimeZone=Europe/Moscow",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = DB.AutoMigrate(&Number{})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/number", handler)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server started")
	n := Number{}
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		panic(err)
	}

	res := []Number{}
	db := DB.Where("number IN ?", []int{n.Number, n.Number + 1}).Find(&res)
	if db.Error != nil {
		panic(err)
	}

	resp := Response{}

	w.Header().Set("Content-Type", "application/json")

	if len(res) == 2 {
		resp.Code = 500
		resp.Message = "Both entries of the number and the number less by one are already in the database"
		w.WriteHeader(http.StatusInternalServerError)
	}

	if len(res) == 1 {

		if res[0].Number == n.Number {

			resp.Code = 500
			resp.Message = "This number are already in the database"
			w.WriteHeader(http.StatusInternalServerError)

		} else {

			resp.Code = 500
			resp.Message = "Number less by one are already in the database"
			w.WriteHeader(http.StatusInternalServerError)

		}
	}

	if db.RowsAffected == 0 {

		db = DB.Create(&n)
		if db.Error != nil {
			panic(err)
		}

		resp.Code = 200
		resp.Message = "Ok"
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(&resp)

	//fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
