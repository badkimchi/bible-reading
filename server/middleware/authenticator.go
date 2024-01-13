package middleware

import (
	"app/util/resp"
	"errors"
	"fmt"
	"github.com/go-chi/jwtauth"
	"net/http"
	"strconv"
	"strings"
)

func Authenticator(requiredLevel int) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if strings.ToUpper(r.Method) != "OPTIONS" {
					err, statusCode := authenticate(r, requiredLevel)
					if statusCode == 401 {
						resp.InvalidAuth(w, r, err)
						return
					}
					if statusCode == 403 {
						resp.Forbidden(w, r, err)
						return
					}
				}
				next.ServeHTTP(w, r)
			},
		)
	}
}

func authenticate(r *http.Request, requiredLevel int) (error, int) {
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return err, 401
	}
	if token == nil {
		return errors.New("no token found"), 401
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	//Refresh token is not allowed for authentication
	if claims["token_type"] != "auth" {
		return errors.New("not a valid auth token"), 401
	}
	userLevelStr := claims["level"].(string)
	userLevel, err := strconv.Atoi(userLevelStr)
	if err != nil {
		return errors.New("unable to extract user level from the token"), 0
	}
	if !hasPermission(userLevel, requiredLevel) {
		msg := fmt.Sprintf(
			"the account level %d is not allowed to use the api end point: %s: %s",
			userLevel, r.URL.Path, r.Method,
		)
		return errors.New(msg), 403
	}

	return nil, 200
}

func hasPermission(userLevel int, requiredLevel int) bool {
	return userLevel >= requiredLevel
}
