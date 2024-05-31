package sendcontroller

import (
	"fmt"
	"net/http"

	"github.com/jovinkendrico/futurefarmerapi/helper"
	"github.com/jovinkendrico/futurefarmerapi/models"
)

func GetRelayStatus(w http.ResponseWriter, r *http.Request) {
	// Retrieve the RelayStatus record by ID
	var relayStatus models.RelayStatus
	if err := models.DB.First(&relayStatus, "1").Error; err != nil {
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}

	var relayConfig models.RelayConfig
	result := models.DB.First(&relayConfig)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Check the status of each field
	statuses := map[string]string{
		"Relay1_is": checkStatus(relayStatus.Ph_up),
		"Relay2_is": checkStatus(relayStatus.Ph_down),
		"Relay3_is": checkStatus(relayStatus.Nut_a),
		"Relay4_is": checkStatus(relayStatus.Nut_b),
		"Relay5_is": checkStatus(relayStatus.Fan),
		"Relay6_is": checkStatus(relayStatus.Light),
		"is_sync":   fmt.Sprintf("%d", relayConfig.IsSync),
	}

	// Respond with the statuses in JSON format
	helper.ResponseJSON(w, http.StatusOK, statuses)
}

func checkStatus(value int64) string {
	if value == 1 {
		return "on"
	}
	return "off"
}
