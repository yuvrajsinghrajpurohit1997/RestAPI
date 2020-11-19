package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	_ "github.com/sparrc/go-ping"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB global variable..
var DB *gorm.DB

type users struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Active    bool `json:"Active" gorm:"default:true"`
	CreatedAt string
}

//NewUser function ..
func NewUser(w http.ResponseWriter, r *http.Request) {
	connect(w, r)
	var user users
	json.NewDecoder(r.Body).Decode(&user)
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

//GetUser function..
func GetUser(w http.ResponseWriter, r *http.Request) {
	connect(w, r)
	var user []users
	DB.Find(&user)
	json.NewEncoder(w).Encode(&user)

}

//GetUserID function..
func GetUserID(w http.ResponseWriter, r *http.Request) {
	connect(w, r)
	vars := mux.Vars(r)
	userID := vars["id"]
	var getUser users
	DB := DB.Where("ID = ?", userID).Find(&getUser)
	DB.Find(&getUser)
	userDetails, _ := &getUser, DB
	res, _ := json.Marshal(&userDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(res)
}

// connect function..
func connect(w http.ResponseWriter, r *http.Request) {
	dsn := "host=localhost user=postgres password=1234qwer dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Can not connect to the database")
	}
	DB = db

}
