package configcontroller

import (
	"net/http"

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
	}