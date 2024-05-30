package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jovinkendrico/futurefarmerapi/controllers/authcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/configcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/dashboardcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/datacontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/sendcontroller"
	"github.com/jovinkendrico/futurefarmerapi/models"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/insertdata", datacontroller.InsertData).Methods("POST")
	r.HandleFunc("/getconfig", configcontroller.GetConfig).Methods("GET")
	r.HandleFunc("/relaystatus/{id}", sendcontroller.GetRelayStatus).Methods("GET")
	r.HandleFunc("/api/v1/dashboard", dashboardcontroller.Index).Methods("GET")

	fmt.Printf("Server is running !!!")
	log.Fatal(http.ListenAndServe(":8000", r))
}
