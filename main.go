package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/jovinkendrico/futurefarmerapi/controllers/authcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/configcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/dashboardcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/datacontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/plantcontroller"
	"github.com/jovinkendrico/futurefarmerapi/controllers/sendcontroller"
	"github.com/jovinkendrico/futurefarmerapi/models"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	models.ConnectDatabase()
	r := mux.NewRouter()

	//IOT API
	r.HandleFunc("/insertdata", datacontroller.InsertData).Methods("POST")
	r.HandleFunc("/getconfig", configcontroller.GetConfig).Methods("GET")
	r.HandleFunc("/updaterelay", configcontroller.UpdateRelay).Methods("POST")
	r.HandleFunc("/relaystatus", sendcontroller.GetRelayStatus).Methods("GET")

	//ANDROID API
	r.HandleFunc("/api/v1/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/api/v1/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/api/v1/logout", authcontroller.Logout).Methods("GET")
	r.HandleFunc("/api/v1/dashboard", dashboardcontroller.Index).Methods("GET")
	r.HandleFunc("/api/v1/updateconfig", configcontroller.UpdateConfig).Methods("PUT")
	r.HandleFunc("/api/v1/updaterelay", configcontroller.UpdateRelayStatus).Methods("PUT")
	r.HandleFunc("/api/v1/plant", plantcontroller.Index).Methods("GET")
	r.HandleFunc("/api/v1/plant", plantcontroller.Insert).Methods("POST")

	fmt.Printf("Server is running !!!")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
