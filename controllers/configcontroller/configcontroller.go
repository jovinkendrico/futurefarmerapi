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

func UpdateConfig(w http.ResponseWriter, r *http.Request) {
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
	var DB *gorm.DB
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

	relay_1 := r.FormValue("Relay_1")
	relay_1_int, err := strconv.ParseInt(relay_1, 36, 64)
	if err != nil {
		http.Error(w, "Relay 1 Error", http.StatusBadRequest)
		return
	}
	if relayStatus.Ph_up == 1 && relay_1_int == 0 {
		relayHistory := models.RelayHistory{
			Type:   "PH UP",
			Status: "OFF",
		}
		result := DB.Create(&relayHistory)
		if result.Error != nil {
			panic("failed to insert relay history record")
		}
	}
	relayStatus.Ph_up = relay_1_int


	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}

	relay_2 := r.FormValue("Relay_2")
	relay_2_int, err := strconv.ParseInt(relay_2, 36, 64)
	if err != nil {
		http.Error(w, "Relay 2 Error", http.StatusBadRequest)
		return
	}
	if relayStatus.Ph_down == 1 && relay_2_int == 0 {
		relayHistory := models.RelayHistory{
			Type:   "PH DOWN",
			Status: "OFF",
		}
		result := DB.Create(&relayHistory)
		if result.Error != nil {
			panic("failed to insert relay history record")
		}
	}
	relayStatus.Ph_down = relay_2_int
	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}

	relay_3 := r.FormValue("Relay_3")
	relay_3_int, err := strconv.ParseInt(relay_3, 36, 64)
	if err != nil {
		http.Error(w, "Relay 3 Error", http.StatusBadRequest)
		return
	}
	if relayStatus.Nut_a == 1 && relay_3_int == 0 {
		relayHistory := models.RelayHistory{
			Type:   "NUTRISI A",
			Status: "OFF",
		}
		result := DB.Create(&relayHistory)
		if result.Error != nil {
			panic("failed to insert relay history record")
		}
	}
	relayStatus.Nut_a = relay_3_int
	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}


	relay_4 := r.FormValue("Relay_4")
	relay_4_int, err := strconv.ParseInt(relay_4, 36, 64)
	if err != nil {
		http.Error(w, "Relay 4 Error", http.StatusBadRequest)
		return
	}
	if relayStatus.Nut_b == 1 && relay_4_int == 0 {
		relayHistory := models.RelayHistory{
			Type:   "NUTRISI B",
			Status: "OFF",
		}
		result := DB.Create(&relayHistory)
		if result.Error != nil {
			panic("failed to insert relay history record")
		}
	}
	relayStatus.Nut_b = relay_4_int
	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}


	relay_5 := r.FormValue("Relay_5")
	relay_5_int, err := strconv.ParseInt(relay_5, 36, 64)
	if err != nil {
		http.Error(w, "Relay 5 Error", http.StatusBadRequest)
		return
	}	
	if relayStatus.Fan == 1 && relay_5_int == 0 {
		relayHistory := models.RelayHistory{
			Type:   "FAN",
			Status: "OFF",
		}
		result := DB.Create(&relayHistory)
		if result.Error != nil {
			panic("failed to insert relay history record")
		}
	}
	relayStatus.Fan = relay_5_int
	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}

	relay_6 := r.FormValue("Relay_6")
	relay_6_int, err := strconv.ParseInt(relay_6, 36, 64)
	if err != nil {
		http.Error(w, "Relay 6 Error", http.StatusBadRequest)
		return
	}
	if relayStatus.Light == 1 && relay_6_int == 0 {
		relayHistory := models.RelayHistory{
			Type:   "LIGHT",
			Status: "OFF",
		}
		result := DB.Create(&relayHistory)
		if result.Error != nil {
			panic("failed to insert relay history record")
		}
	}
	relayStatus.Light = relay_6_int
	if saveResult := models.DB.Save(&relayStatus); saveResult.Error != nil {
		http.Error(w, saveResult.Error.Error(), http.StatusInternalServerError)
	}

}
