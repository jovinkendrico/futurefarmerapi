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
	"github.com/jovinkendrico/futurefarmerapi/middlewares"
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
	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")

	//ANDROID API
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/v1/logout", authcontroller.Logout).Methods("GET")
	api.HandleFunc("/v1/dashboard", dashboardcontroller.Index).Methods("GET")

	//Relay config
	api.HandleFunc("/v1/getrelayconfig", configcontroller.GetRelayConfig).Methods("GET")
	api.HandleFunc("/v1/updaterelayconfig", configcontroller.UpdateRelayConfig).Methods("PUT")

	//Level Config
	api.HandleFunc("/v1/getlevelconfig", configcontroller.GetLevelConfig).Methods("GET")
	api.HandleFunc("/v1/updatelevelconfig", configcontroller.UpdateLevelConfig).Methods("PUT")

	//relay status on off manual auto
	api.HandleFunc("/v1/updaterelay", configcontroller.UpdateRelayStatus).Methods("PUT")

	//tanaman
	api.HandleFunc("/v1/plant", plantcontroller.Index).Methods("GET")
	api.HandleFunc("/v1/plant", plantcontroller.Insert).Methods("POST")
	//use middleware jwt for android

	api.Use(middlewares.JWTMiddleware)
	fmt.Printf("Server is running !!!")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
