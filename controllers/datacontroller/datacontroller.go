package datacontroller

import (
	"net/http"
	"strconv"
	"time"

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
		http.Error(w, "Invalid temperature value", http.StatusBadRequest)
		return
	}

	// Extract and validate humidity
	humidityStr := r.FormValue("humidity")
	humidity, err := strconv.ParseFloat(humidityStr, 64)
	if err != nil {
		http.Error(w, "Invalid humidity value", http.StatusBadRequest)
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

	now := time.Now()
	checkInterval := -3 * time.Hour

	var phUpHistory models.RelayHistory
	result = models.DB.Where("type = ?", "PH UP").Order("created_at DESC").First(&phUpHistory)

	if result.Error != nil || phUpHistory.CreatedAt.Before(now.Add(checkInterval)) {
		if ph < levelConfig.Ph_low && relayStatus.Is_manual_1 == 0 {
			relayStatus.Ph_up = 1
			var relayHistory models.RelayHistory
			relayHistory.Type = "PH UP"
			relayHistory.Status = "ON"
			err := models.DB.Create(&relayHistory).Error
			if err != nil {
				panic("failed to insert relay history record")
			}
		}
	}

	var phDownHistory models.RelayHistory
	result = models.DB.Where("type = ?", "PH DOWN").Order("created_at DESC").First(&phDownHistory)
	if result.Error != nil || phDownHistory.CreatedAt.Before(now.Add(checkInterval)) {
		if ph > levelConfig.Ph_high && relayStatus.Is_manual_2 == 0 {
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
	}

	if tds < levelConfig.Tds {
		var nutAHistory models.RelayHistory
		result = models.DB.Where("type = ?", "NUTRISI A").Order("created_at DESC").First(&nutAHistory)
		if result.Error != nil || nutAHistory.CreatedAt.Before(now.Add(checkInterval)) {
			if relayStatus.Is_manual_3 == 0 {
				relayStatus.Nut_a = 1

				relayHistory := models.RelayHistory{
					Type:   "NUTRISI A",
					Status: "ON",
				}
				result := models.DB.Create(&relayHistory).Error
				if result != nil {
					panic("failed to insert relay history record")
				}
			}
		}

		var nutBHistory models.RelayHistory
		result = models.DB.Where("type = ?", "NUTRISI B").Order("created_at DESC").First(&nutBHistory)
		if result.Error != nil || nutBHistory.CreatedAt.Before(now.Add(checkInterval)) {
			if relayStatus.Is_manual_4 == 0 {
				relayStatus.Nut_b = 1
				relayHistory_2 := models.RelayHistory{
					Type:   "NUTRISI B",
					Status: "ON",
				}
				result_2 := models.DB.Create(&relayHistory_2).Error
				if result_2 != nil {
					panic("failed to insert relay history record")
				}
			}
		}
	}

	if (temperature > levelConfig.Temperature_high || humidity > levelConfig.Humidity) && relayStatus.Is_manual_5 == 0 {
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

	// if temperature > levelConfig.Temperature_high  && relayStatus.Is_manual_6 == 0  {
	// 	relayStatus.Light = 1
	// 	relayHistory := models.RelayHistory{
	// 		Type:   "LIGHT",
	// 		Status: "ON",
	// 	}
	// 	result := models.DB.Create(&relayHistory).Error
	// 	if result != nil {
	// 		panic("failed to insert relay history record")
	// 	}
	// }

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
