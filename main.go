package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jovinkendrico/futurefarmerapi/controllers/authcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/dashboardcontroller"
	"github.com/jovinkendrico/futurefarmerapi/models"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	r.HandleFunc("/api/v1/dashboard", dashboardcontroller.Index).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
