package base

import (
	"app/internal/routes/base/health"
	"app/internal/routes/base/root"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Get("/", root.Handler)

		r.Get("/health", health.Handler)
	})
}
