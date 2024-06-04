package dashboardcontroller

import (
	"net/http"

	"github.com/jovinkendrico/futurefarmerapi/helper"
	"github.com/jovinkendrico/futurefarmerapi/models"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var SensorData models.SensorData
	if err := models.DB.Last(&SensorData).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"error": "true", "message": "Record not found"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"error": "true", "message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}

	}
	data := map[string]interface{}{
		"error":      "false",
		"message":    "Record found",
		"id":         SensorData.Id,
		"ph":         SensorData.Ph,
		"tds":        SensorData.Tds,
		"suhu":       SensorData.Temperature,
		"kelembapan": SensorData.Humidity,
		"created_at": SensorData.CreatedAt,
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
