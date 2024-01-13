package main

import (
	"app/conf"
	"app/middleware"
	"app/sql/db"
	"github.com/dillonstreator/opentelemetry-go-contrib/instrumentation/net/http/otelhttp"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"log/slog"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func defineRoutes(mux *chi.Mux, cfg *conf.Config, logger *slog.Logger, token *jwtauth.JWTAuth, queries *db.Queries) {
	setCommonMiddleware(mux, logger)
	cont, err := controllers(cfg, token, queries)
	if err != nil {
		panic(err)
	}
	mux.Get(cfg.HealthEndpoint, handleHealthCheck)
	defineStaticRoutes(mux)
	defineAPIRoutes(cont, mux, token)
}

func setCommonMiddleware(mux *chi.Mux, logger *slog.Logger) {
	mux.Use(chimiddleware.Recoverer)
	mux.Use(middleware.TrustProxy(logger))
	mux.Use(otelhttp.NewMiddleware("chi"))
	mux.Use(middleware.RequestLogger(logger))
	mux.Use(middleware.CorsHeaders())
}

func defineStaticRoutes(mux *chi.Mux) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	workDir := filepath.Dir(ex)
	feDir := "web/dist"
	feBasePath := path.Join(workDir, feDir)
	indexFilePath := path.Join(feBasePath, "index.html")
	mux.Get("/assets/*", func(w http.ResponseWriter, r *http.Request) {
		fullFilePath := path.Join(feBasePath, r.URL.Path)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, fullFilePath)
	})
	mux.Get("/", serveFileHandler(indexFilePath))
	mux.NotFound(serveFileHandler(indexFilePath))
}

func defineAPIRoutes(cont reqControllers, mux *chi.Mux, token *jwtauth.JWTAuth) {
	//public path
	mux.Post("/auth/login", cont.AuthC.Login)

	mux.Route("/api", func(api chi.Router) {
		api.Use(middleware.Timeout(time.Second * 10))
		api.Use(jwtauth.Verifier(token))
		api.Use(middleware.Authenticator(0))

		api.Get("/accounts/{id}", cont.AccC.GetAccount)
	})
}

func handleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func serveFileHandler(fPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		http.ServeFile(w, r, fPath)
	}
}
