package sendcontroller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jovinkendrico/futurefarmerapi/helper"
	"github.com/jovinkendrico/futurefarmerapi/models"
)

func GetRelayStatus(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL parameters
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	// Convert the ID to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Retrieve the RelayStatus record by ID
	var relayStatus models.RelayStatus
	if err := models.DB.First(&relayStatus, id).Error; err != nil {
		http.Error(w, "Record not found", http.StatusNotFound)
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
	}

	// Respond with the statuses in JSON format
	helper.ResponseJSON(w, http.StatusOK, statuses)
}

func checkStatus(value float64) string {
	if value == 1 {
		return "on"
	}
	return "off"
}
