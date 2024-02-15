package resp

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

type ResponseObj struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func OK(w http.ResponseWriter, r *http.Request, msg string) {
	render.JSON(w, r, ResponseObj{Message: msg, Data: msg})
}

func Data(w http.ResponseWriter, r *http.Request, v interface{}) {
	render.JSON(w, r, ResponseObj{Message: "ok", Data: v})
}

func Bad(w http.ResponseWriter, r *http.Request, err error) {
	if err == nil {
		w.WriteHeader(500)
		render.JSON(w, r, ResponseObj{Message: "unable to parse error", Data: "unable to parse error"})
	}

	statusCode := 400
	if strings.Contains(err.Error(), "record not found") {
		statusCode = 404
	}

	w.WriteHeader(statusCode)
	render.JSON(w, r, ResponseObj{Message: err.Error(), Data: err.Error()})
}
func InvalidAuth(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(401)
	msg := fmt.Sprintf("invalid token: %v", err)
	render.JSON(w, r, ResponseObj{Message: msg, Data: msg})
}

func Forbidden(w http.ResponseWriter, r *http.Request, err error) {
	w.WriteHeader(403)
	render.JSON(w, r, ResponseObj{Message: err.Error()})
}
