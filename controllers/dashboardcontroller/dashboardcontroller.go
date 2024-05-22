package dashboardcontroller

import (
	"net/http"
	"time"

	"github.com/jovinkendrico/futurefarmerapi/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]interface{}{
		{
			"id":         1,
			"datetime":   time.Now(),
			"ph":         6.4,
			"tds":        25,
			"suhu":       36,
			"kelembapan": 21,
		},
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
