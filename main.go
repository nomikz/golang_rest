package main

import (
	"./controllers"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// here
	router := mux.NewRouter()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	router.Handle("/", staticC.Home).Methods("GET")
	router.Handle("/contacts", staticC.Contacts).Methods("GET")
	router.HandleFunc("/sign-up", usersC.New).Methods("GET")
	router.HandleFunc("/sign-up", usersC.Create).Methods("POST")

	_ = http.ListenAndServe(":3000", router)
}