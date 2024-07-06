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
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}
	models.ConnectDatabase()
	r := mux.NewRouter()

	c := cron.New()

	_, err = c.AddFunc("CRON_TZ=Asia/Jakarta */30 * * * *", func() {
		var RelayStatus models.RelayStatus
		if err := models.DB.First(&RelayStatus).Error; err != nil {
			return
		}
		if RelayStatus.Light == 0 {
			RelayStatus.Light = 1
		} else {
			RelayStatus.Light = 0
		}
		if err := models.DB.Save(&RelayStatus).Error; err != nil {
			return
		}
		fmt.Println("Turning database 'on' at odd hours")
	})
	if err != nil {
		log.Fatalf("Error adding cron job: %v", err)
	}

	// Cron job to turn database flag 'off' at every even hour
	// _, err = c.AddFunc("CRON_TZ=Asia/Jakarta 0 0-22/2 * * *", func() {
	// 	var RelayStatus models.RelayStatus
	// 	if err := models.DB.First(&RelayStatus).Error; err != nil {
	// 		return
	// 	}
	// 	RelayStatus.Light = 0
	// 	if err := models.DB.Save(&RelayStatus).Error; err != nil {
	// 		return
	// 	}
	// 	fmt.Println("Turning database 'off' at even hours")
	// })
	// if err != nil {
	// 	log.Fatalf("Error adding cron job: %v", err)
	// }

	c.Start()
	defer c.Stop()

	iotAPI := r.PathPrefix("/iot").Subrouter()
	iotAPI.Use(middlewares.APIKEYMiddleware)

	iotAPI.HandleFunc("/insertdata", datacontroller.InsertData).Methods("POST")
	iotAPI.HandleFunc("/getconfig", configcontroller.GetConfig).Methods("GET")
	iotAPI.HandleFunc("/updaterelay", configcontroller.UpdateRelay).Methods("POST")
	iotAPI.HandleFunc("/relaystatus", sendcontroller.GetRelayStatus).Methods("GET")

	//login logic
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
	api.HandleFunc("/v1/getrelaystatus", configcontroller.GetRelayStatus).Methods("GET")
	api.HandleFunc("/v1/updaterelayphup", configcontroller.UpdateRelayPhUp).Methods("PATCH")
	api.HandleFunc("/v1/updaterelayphdown", configcontroller.UpdateRelayPhDown).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaynuta", configcontroller.UpdateRelayNutA).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaynutb", configcontroller.UpdateRelayNutB).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaynutrisi", configcontroller.UpdateRelayNutrisi).Methods("PATCH")
	api.HandleFunc("/v1/updaterelayfan", configcontroller.UpdateRelayFan).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaylight", configcontroller.UpdateRelayLight).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaymanualone", configcontroller.UpdateRelayManualOne).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaymanualtwo", configcontroller.UpdateRelayManualTwo).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaymanualthree", configcontroller.UpdateRelayManualThree).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaymanualfour", configcontroller.UpdateRelayManualFour).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaymanualfive", configcontroller.UpdateRelayManualFive).Methods("PATCH")
	api.HandleFunc("/v1/updaterelaymanualsix", configcontroller.UpdateRelayManualSix).Methods("PATCH")

	//tanaman
	api.HandleFunc("/v1/plant", plantcontroller.Index).Methods("GET")
	api.HandleFunc("/v1/plant", plantcontroller.Insert).Methods("POST")

	//get relay history
	api.HandleFunc("/v1/getrelayhistory", datacontroller.GetRelayHistory).Methods("GET")
	//use middleware jwt for android

	api.Use(middlewares.JWTMiddleware)
	fmt.Printf("Server is running !!!")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
