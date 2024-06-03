package datacontroller

import (
	"net/http"
	"strconv"

	"github.com/jovinkendrico/futurefarmerapi/helper"
	"github.com/jovinkendrico/futurefarmerapi/models"
)

func InsertData(w http.ResponseWriter, r *http.Request) {

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Extract and validate ph
	phStr := r.FormValue("ph")
	ph, err := strconv.ParseFloat(phStr, 64)
	if err != nil {
		http.Error(w, "Invalid ph value", http.StatusBadRequest)
		return
	}

	// Extract and validate tds
	tdsStr := r.FormValue("tds")
	tds, err := strconv.ParseFloat(tdsStr, 64)
	if err != nil {
		http.Error(w, "Invalid tds value", http.StatusBadRequest)
		return
	}

	// Extract and validate temperature
	temperatureStr := r.FormValue("temperature")
	temperature, err := strconv.ParseFloat(temperatureStr, 64)
	if err != nil {
		http.Error(w, "Invalid temp value", http.StatusBadRequest)
		return
	}

	// Extract and validate humidity
	humidityStr := r.FormValue("humidity")
	humidity, err := strconv.ParseFloat(humidityStr, 64)
	if err != nil {
		http.Error(w, "Invalid humid value", http.StatusBadRequest)
		return
	}

	// Create a new SensorData instance
	data := models.SensorData{
		Ph:          ph,
		Tds:         tds,
		Temperature: temperature,
		Humidity:    humidity,
	}

	var levelConfig models.LevelConfig
	result := models.DB.First(&levelConfig)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	var relayStatus models.RelayStatus
	result_2 := models.DB.First(&relayStatus)
	if result_2.Error != nil {
		http.Error(w, result_2.Error.Error(), http.StatusInternalServerError)
		return
	}


	if ph < levelConfig.Ph_low {
		relayStatus.Ph_up = 1
		var relayHistory models.RelayHistory
		relayHistory.Type = "PH UP"
		relayHistory.Status = "on"
		err := models.DB.Create(&relayHistory).Error
		if  err != nil {
			panic("failed to insert relay history record")
		}
	}

	if ph > levelConfig.Ph_high {
		relayStatus.Ph_down = 1
		relayHistory := models.RelayHistory{
			Type:   "PH DOWN",
			Status: "ON",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
	}
	if tds < levelConfig.Tds {
		relayStatus.Nut_a = 1
		relayStatus.Nut_b = 1
		relayHistory := models.RelayHistory{
			Type:   "NUTRISI A",
			Status: "ON",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayHistory_2 := models.RelayHistory{
			Type:   "NUTRISI B",
			Status: "ON",
		}
		result_2 := models.DB.Create(&relayHistory_2).Error
		if result_2 != nil {
			panic("failed to insert relay history record")
		}

	}

	if temperature < levelConfig.Temperature_low || humidity < levelConfig.Humidity {
		relayStatus.Fan = 1
		relayHistory := models.RelayHistory{
			Type:   "FAN",
			Status: "ON",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
	}

	if temperature > levelConfig.Temperature_high {
		relayStatus.Light = 1
		relayHistory := models.RelayHistory{
			Type:   "LIGHT",
			Status: "ON",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
	}

	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}

	// Insert the data into the database
	if err := models.DB.Create(&data).Error; err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	// Respond with the inserted data
	helper.ResponseJSON(w, http.StatusOK, data)

}
