package main

import (
	"./controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	router.Handle("/", staticC.HomeView).Methods("GET")
	router.Handle("/contacts", staticC.ContactsView).Methods("GET")
	router.HandleFunc("/sign-up", usersC.New).Methods("GET")
	router.HandleFunc("/sign-up", usersC.Create).Methods("POST")

	_ = http.ListenAndServe(":3000", router)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
