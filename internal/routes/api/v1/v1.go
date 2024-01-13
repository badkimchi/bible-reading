package v1

import (
	"app/internal/routes/api/v1/forecast"
	"app/internal/routes/api/v1/temperature"

	"github.com/go-chi/chi/v5"
)

func RegisterPrefix(r chi.Router) {
	r.Route("/v1/", func(r chi.Router) {
		registerRoutes(r)
	})
}

func registerRoutes(r chi.Router) {
	r.Get("/temperature", temperature.Handler)

	r.Get("/forecast/{forecastPeriod}", forecast.Handler)
}
