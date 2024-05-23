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
		http.Error(w, "Invalid temperature value", http.StatusBadRequest)
		return
	}

	// Extract and validate tds
	tdsStr := r.FormValue("tds")
	tds, err := strconv.ParseFloat(tdsStr, 64)
	if err != nil {
		http.Error(w, "Invalid temperature value", http.StatusBadRequest)
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
		Ph:       ph,
		Tds:       tds,
		Temperature:       temperature,
		Humidity: humidity,
	}

	// Insert the data into the database
	if err := models.DB.Create(&data).Error; err != nil {
		http.Error(w, "Failed to insert data", http.StatusInternalServerError)
		return
	}

	// Respond with the inserted data
	helper.ResponseJSON(w, http.StatusOK, data)

}
