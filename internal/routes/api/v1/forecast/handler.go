package forecast

import (
	"app/internal/logger"
	"app/internal/responder"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// mock api response
func Handler(w http.ResponseWriter, r *http.Request) {
	var forecastPeriod = chi.URLParam(r, "forecastPeriod")

	if forecastPeriod == "" {
		http.Error(w, "No forecast forecast specified", http.StatusBadRequest)
		return
	}

	msg := []map[string]string{}

	switch forecastPeriod {
	case "3day":
		msg = append(msg, map[string]string{"temp_c": "18.8", "wind_speed_km": "10", "humidity_percent": "70"})
		msg = append(msg, map[string]string{"temp_c": "25.2", "wind_speed_km": "12", "humidity_percent": "74"})
		msg = append(msg, map[string]string{"temp_c": "27.7", "wind_speed_km": "9", "humidity_percent": "86"})
	default:
		http.Error(w, fmt.Sprintf("unsupported forecast period: %s", forecastPeriod), http.StatusBadRequest)
		return
	}

	if err := responder.JSONPretty(w, msg, http.StatusOK); err != nil {
		logger.StderrWithSource.Error(err.Error())
	}
}
