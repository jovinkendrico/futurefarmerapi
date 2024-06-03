package plantcontroller

import (
	"encoding/json"
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
		"id":         Plant.Id,
		"nama":       Plant.Nama,
		"tanggal":    Plant.Tanggal,
		"umur":       Plant.Umur,
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
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	tanggal, err := time.Parse("2006-01-02", PlantInput.Tanggal)
	if err != nil {
		http.Error(w, "Invalid tanggal format", http.StatusBadRequest)
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
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
