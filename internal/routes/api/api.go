package api

import (
	v1 "app/internal/routes/api/v1"

	"github.com/go-chi/chi/v5"
)

func RegisterPrefix(r chi.Router) {
	r.Route("/api/", func(r chi.Router) {
		v1.RegisterPrefix(r)
	})
}
