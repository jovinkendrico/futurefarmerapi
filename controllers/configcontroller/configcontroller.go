package configcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jovinkendrico/futurefarmerapi/helper"
	"github.com/jovinkendrico/futurefarmerapi/models"
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

func UpdateConfig(w http.ResponseWriter, r *http.Request) {
	var relayConfigInput models.RelayConfig
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&relayConfigInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
	}

	defer r.Body.Close()

	if err := models.DB.Update("1", &relayConfigInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
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
	}

	defer r.Body.Close()

	if err := models.DB.Update("1", &relayStatusInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
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


	if (relayStatus.Ph_up == 1 && relay_id_int == 1 && status=="off") {
		relayHistory := models.RelayHistory{
			Type:   "PH UP",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Ph_up = 0
	} else if (relayStatus.Ph_down == 1 && relay_id_int == 2 && status=="off") {
		relayHistory := models.RelayHistory{
			Type:   "PH DOWN",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Ph_down = 0
	} else if (relayStatus.Nut_a == 1 && relay_id_int == 3 && status=="off") {
		relayHistory := models.RelayHistory{
			Type:   "NUTRISI A",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Nut_a = 0
	} else if (relayStatus.Nut_b == 1 && relay_id_int == 4 && status=="off") {
		relayHistory := models.RelayHistory{
			Type:   "NUTRISI B",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Nut_b = 0
	} else if (relayStatus.Fan == 1 && relay_id_int == 5 && status=="off") {
		relayHistory := models.RelayHistory{
			Type:   "FAN",
			Status: "OFF",
		}
		result := models.DB.Create(&relayHistory).Error
		if result != nil {
			panic("failed to insert relay history record")
		}
		relayStatus.Fan = 0
	} else if (relayStatus.Fan == 1 && relay_id_int == 6 && status=="off") {
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
