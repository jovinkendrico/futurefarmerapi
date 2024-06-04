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
func GetRelayStatus(w http.ResponseWriter, r *http.Request) {
	var RelayStatus models.RelayStatus
	if err := models.DB.Last(&RelayStatus).Error; err != nil {
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
		"id":          RelayStatus.Id,
		"ph_up":       RelayStatus.Ph_up,
		"is_manual_1": RelayStatus.Is_manual_1,
		"ph_down":     RelayStatus.Ph_down,
		"is_manual_2": RelayStatus.Is_manual_2,
		"nut_a":       RelayStatus.Nut_a,
		"is_manual_3": RelayStatus.Is_manual_3,
		"nut_b":       RelayStatus.Nut_b,
		"is_manual_4": RelayStatus.Is_manual_4,
		"light":       RelayStatus.Light,
		"is_manual_5": RelayStatus.Is_manual_5,
		"fan":         RelayStatus.Fan,
		"is_manual_6": RelayStatus.Is_manual_6,
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
func UpdateRelayPhUp(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Ph_up int64 `json:"ph_up"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Ph_up != 0 && RelayInput.Ph_up != 1 {
		response := map[string]string{"message": "Ph Up value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Ph_up = RelayInput.Ph_up
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Ph Up updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayPhDown(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Ph_down int64 `json:"ph_down"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Ph_down != 0 && RelayInput.Ph_down != 1 {
		response := map[string]string{"message": "Ph Down value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Ph_down = RelayInput.Ph_down
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Ph Down updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayNutA(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Nut_a int64 `json:"nut_a"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Nut_a != 0 && RelayInput.Nut_a != 1 {
		response := map[string]string{"message": "Nut A value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Nut_a = RelayInput.Nut_a
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Nut A updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayNutB(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Nut_b int64 `json:"nut_b"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Nut_b != 0 && RelayInput.Nut_b != 1 {
		response := map[string]string{"message": "Nut B value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Nut_b = RelayInput.Nut_b
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Nut B updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayFan(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Fan int64 `json:"fan"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Fan != 0 && RelayInput.Fan != 1 {
		response := map[string]string{"message": "Fan value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Fan = RelayInput.Fan
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Fan updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayLight(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Light int64 `json:"light"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Light != 0 && RelayInput.Light != 1 {
		response := map[string]string{"message": "Light value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Light = RelayInput.Light
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Light updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayManualOne(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Manual_One int64 `json:"manual_one"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Manual_One != 0 && RelayInput.Manual_One != 1 {
		response := map[string]string{"message": "Manual One value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Is_manual_1 = RelayInput.Manual_One
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Manual One updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayManualTwo(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Manual_Two int64 `json:"manual_two"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Manual_Two != 0 && RelayInput.Manual_Two != 1 {
		response := map[string]string{"message": "Manual Two value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Is_manual_2 = RelayInput.Manual_Two
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Manual Two updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayManualThree(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Manual_Three int64 `json:"manual_three"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Manual_Three != 0 && RelayInput.Manual_Three != 1 {
		response := map[string]string{"message": "Manual Three value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Is_manual_3 = RelayInput.Manual_Three
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Manual Three updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayManualFour(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Manual_Four int64 `json:"manual_four"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Manual_Four != 0 && RelayInput.Manual_Four != 1 {
		response := map[string]string{"message": "Manual Four value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Is_manual_4 = RelayInput.Manual_Four
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Manual Four updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayManualFive(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Manual_Five int64 `json:"manual_five"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Manual_Five != 0 && RelayInput.Manual_Five != 1 {
		response := map[string]string{"message": "Manual Five value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Is_manual_5 = RelayInput.Manual_Five
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Manual Five updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
func UpdateRelayManualSix(w http.ResponseWriter, r *http.Request) {
	var RelayInput struct {
		Manual_Six int64 `json:"manual_six"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&RelayInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	if RelayInput.Manual_Six != 0 && RelayInput.Manual_Six != 1 {
		response := map[string]string{"message": "Manual Six value must be 0 or 1"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var RelayStatus models.RelayStatus
	if err := models.DB.First(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	RelayStatus.Is_manual_6 = RelayInput.Manual_Six
	if err := models.DB.Save(&RelayStatus).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	// Respond with success message
	response := map[string]string{"message": "Manual Six updated successfully"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

// func UpdateRelayStatus(w http.ResponseWriter, r *http.Request) {
// 	var relayStatusInput models.RelayStatus
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&relayStatusInput); err != nil {
// 		response := map[string]string{"message": err.Error()}
// 		helper.ResponseJSON(w, http.StatusBadRequest, response)
// 		return
// 	}

// 	defer r.Body.Close()

// 	if err := models.DB.Update("1", &relayStatusInput).Error; err != nil {
// 		response := map[string]string{"message": err.Error()}
// 		helper.ResponseJSON(w, http.StatusInternalServerError, response)
// 		return
// 	}
// 	response := map[string]string{"message": "success"}
// 	helper.ResponseJSON(w, http.StatusOK, response)
// }

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
