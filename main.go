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

	r.HandleFunc("/api/v1/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/api/v1/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/api/v1/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/api/v1/insertdata", datacontroller.InsertData).Methods("POST")
	r.HandleFunc("/api/v1/getconfig", configcontroller.GetConfig).Methods("GET")
	r.HandleFunc("/api/v1/updaterelay", configcontroller.UpdateRelay).Methods("POST")
	r.HandleFunc("/api/v1/relaystatus/{id}", sendcontroller.GetRelayStatus).Methods("GET")
	r.HandleFunc("/api/v1/dashboard", dashboardcontroller.Index).Methods("GET")

	fmt.Printf("Server is running !!!")
	log.Fatal(http.ListenAndServe(":8000", r))
}
