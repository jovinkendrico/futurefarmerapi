package plantcontroller

import (
	"encoding/json"
	"net/http"

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

	var PlantInput models.Plant
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&PlantInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	defer r.Body.Close()

	if err := models.DB.Create(&PlantInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
