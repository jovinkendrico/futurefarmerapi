package configcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jovinkendrico/futurefarmerapi/helper"
	"github.com/jovinkendrico/futurefarmerapi/models"
	"gorm.io/gorm"
)

func GetConfig(w http.ResponseWriter, r *http.Request) {
	var relayConfig models.RelayConfig
	result := models.DB.First(&relayConfig)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "No records found", http.StatusNotFound)
		return
	}

	helper.ResponseJSON(w, http.StatusOK, relayConfig)
	relayConfig.IsSync = 1
	if saveResult := models.DB.Save(&relayConfig); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
		return
	}
}

func GetRelayConfig(w http.ResponseWriter, r *http.Request) {
	var RelayConfig models.RelayConfig
	if err := models.DB.Last(&RelayConfig).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Record not found"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}
	data := map[string]interface{}{
		"id":      RelayConfig.Id,
		"ph_up":   RelayConfig.Ph_up,
		"ph_down": RelayConfig.Ph_down,
		"nut_a":   RelayConfig.Nut_A,
		"nut_b":   RelayConfig.Nut_B,
		"fan":     RelayConfig.Fan,
		"light":   RelayConfig.Light,
		"is_sync": RelayConfig.IsSync,
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}

func UpdateRelayConfig(w http.ResponseWriter, r *http.Request) {
	var relayConfigInput models.RelayConfig
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&relayConfigInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Update("1", &relayConfigInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func GetLevelConfig(w http.ResponseWriter, r *http.Request) {
	var LevelConfig models.LevelConfig
	if err := models.DB.Last(&LevelConfig).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Record not found"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}
	data := map[string]interface{}{
		"id":               LevelConfig.Id,
		"ph_high":          LevelConfig.Ph_high,
		"ph_low":           LevelConfig.Ph_low,
		"tds":              LevelConfig.Tds,
		"temperature_low":  LevelConfig.Temperature_low,
		"temperature_high": LevelConfig.Temperature_high,
		"humidity":         LevelConfig.Humidity,
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}

func UpdateLevelConfig(w http.ResponseWriter, r *http.Request) {
	var LevelConfigInput models.LevelConfig
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&LevelConfigInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Update("1", &LevelConfigInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func UpdateRelayStatus(w http.ResponseWriter, r *http.Request) {
	var relayStatusInput models.RelayStatus
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&relayStatusInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Update("1", &relayStatusInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func UpdateRelay(w http.ResponseWriter, r *http.Request) {
	var relayStatus models.RelayStatus
	result := models.DB.First(&relayStatus)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	relay_id := r.FormValue("relay_id")
	relay_id_int, err := strconv.ParseInt(relay_id, 36, 64)
	if err != nil {
		http.Error(w, "Relay Id Error", http.StatusBadRequest)
		return
	}
	status := r.FormValue("status")

	if relayStatus.Ph_up == 1 && relay_id_int == 1 && status == "off" {
		relayHistory := models.RelayHistory{
			Type:   "PH UP",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Ph_up = 0
	} else if relayStatus.Ph_down == 1 && relay_id_int == 2 && status == "off" {
		relayHistory := models.RelayHistory{
			Type:   "PH DOWN",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Ph_down = 0
	} else if relayStatus.Nut_a == 1 && relay_id_int == 3 && status == "off" {
		relayHistory := models.RelayHistory{
			Type:   "NUTRISI A",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Nut_a = 0
	} else if relayStatus.Nut_b == 1 && relay_id_int == 4 && status == "off" {
		relayHistory := models.RelayHistory{
			Type:   "NUTRISI B",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Nut_b = 0
	} else if relayStatus.Fan == 1 && relay_id_int == 5 && status == "off" {
		relayHistory := models.RelayHistory{
			Type:   "FAN",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Fan = 0
	} else if relayStatus.Fan == 1 && relay_id_int == 6 && status == "off" {
		relayHistory := models.RelayHistory{
			Type:   "LIGHT",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Light = 0
	}

	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}
}
