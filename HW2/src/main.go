package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"sync"
	"strconv"

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

func Finishing() {
	fmt.Println("Finishing")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGKILL)
	signal.Notify(c, os.Interrupt, syscall.SIGSTOP)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		Finishing()
		os.Exit(1)
	}()

	fmt.Println(os.Environ())

	host := os.Getenv("DB_HOST_IP")
	port := os.Getenv("DB_HOST_PORT")
	var err error
	portI, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	user := os.Getenv("USER_DB_NAME")
	password := os.Getenv("USER_DB_PASSWORD")
	db := os.Getenv("DB_NAME")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		host, user,
		password, db,
		portI, timezone)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		for i := 0; i < 1000; i++ {
			time.Sleep(5 * time.Second)
			DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil{
				break
			}
			
		}
		log.Println(err)
	}

	err = DB.AutoMigrate(&Number{})
	if err != nil {
		log.Fatal(err)
	}
	mu=sync.Mutex{}
	http.HandleFunc("/number", handler)

	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var mu sync.Mutex
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	n := Number{}
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println("accept request with number: " + strconv.Itoa(n.Number))
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
		niceResp := make(map[string]interface{})
		niceResp["number"] = n.Number + 1
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&niceResp)
		return
	}

	json.NewEncoder(w).Encode(&resp)

}


/*

docker build -t iot_cnd_hw2 .
docker image tag iot_cnd_hw2:latest farank338/iot_cnd_hw2:latest
docker image push farank338/iot_cnd_hw2:latest
*/