package plantcontroller

import (
	"encoding/json"
	"math"
	"net/http"
	"time"

	"github.com/jovinkendrico/futurefarmerapi/helper"
	"github.com/jovinkendrico/futurefarmerapi/models"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var Plant models.Plant
	if err := models.DB.Last(&Plant).Error; err != nil {
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
	currentDate := time.Now()
	tanggalDb := Plant.Tanggal
	umurCalculated := math.Round(currentDate.Sub(tanggalDb).Hours() / 24) // Umur in days rounded to nearest int

	// Calculate the Panen
	panen := math.Round(Plant.Umur - umurCalculated)

	data := map[string]interface{}{
		"error":      "false",
		"message":    "Record found",
		"id":         Plant.Id,
		"nama":       Plant.Nama,
		"tanggal":    Plant.Tanggal,
		"umur":       int(umurCalculated),
		"panen":      int(panen),
		"created_at": Plant.CreatedAt,
		"updated_at": Plant.UpdatedAt,
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	var PlantInput struct {
		Nama    string  `json:"nama"`
		Tanggal string  `json:"tanggal"`
		Umur    float64 `json:"umur"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&PlantInput); err != nil {
		response := map[string]string{"error": "true", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	tanggal, err := time.Parse("2006-01-02", PlantInput.Tanggal)
	if err != nil {
		response := map[string]string{"error": "true", "message": "Invalid date format"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	plant := models.Plant{
		Nama:      PlantInput.Nama,
		Tanggal:   tanggal,
		Umur:      PlantInput.Umur,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	defer r.Body.Close()

	if err := models.DB.Create(&plant).Error; err != nil {
		response := map[string]string{"error": "true", "message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"error": "false", "message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
