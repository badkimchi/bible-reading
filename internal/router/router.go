package router

import (
	internalMiddleware "app/internal/middleware"
	"app/internal/routes/api"
	"app/internal/routes/base"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var ChiRouter = chi.NewRouter()

func init() {
	ChiRouter.Use(middleware.Recoverer)
	ChiRouter.Use(internalMiddleware.TrustProxy(internalMiddleware.PrivateRanges()))
	ChiRouter.Use(internalMiddleware.Logger())

	base.RegisterRoutes(ChiRouter)

	api.RegisterPrefix(ChiRouter)
}
