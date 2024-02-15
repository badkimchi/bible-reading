package middleware

import (
	"app/util/resp"
	"context"
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
				var claims map[string]interface{}
				if strings.ToUpper(r.Method) != "OPTIONS" {
					var err error
					var statusCode int
					err, statusCode, claims = authenticate(r, requiredLevel)
					if statusCode == 401 {
						resp.InvalidAuth(w, r, err)
						return
					}
					if statusCode == 403 {
						resp.Forbidden(w, r, err)
						return
					}
				}
				newReq := storeDecodedInfoInContext(claims, r)
				next.ServeHTTP(w, newReq)
			},
		)
	}
}

func storeDecodedInfoInContext(claims map[string]interface{}, r *http.Request) *http.Request {
	ctx := r.Context()
	userId, exists := claims["user_id"]
	if exists {
		ctx = context.WithValue(ctx, "user_id", userId.(string))
	}
	level, exists := claims["level"]
	if exists {
		ctx = context.WithValue(ctx, "level", level.(string))
	}
	return r.WithContext(ctx)
}

func authenticate(r *http.Request, requiredLevel int) (error, int, map[string]interface{}) {
	token, _, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return err, 401, nil
	}
	if token == nil {
		return errors.New("no token found"), 401, nil
	}

	_, claims, _ := jwtauth.FromContext(r.Context())
	//Refresh token is not allowed for authentication
	if claims["token_type"] != "auth" {
		return errors.New("not a valid auth token"), 401, nil
	}
	userLevelStr := claims["level"].(string)
	userLevel, err := strconv.Atoi(userLevelStr)
	if err != nil {
		return errors.New("unable to extract user level from the token"), 0, nil
	}
	if !hasPermission(userLevel, requiredLevel) {
		msg := fmt.Sprintf(
			"the account level %d is not allowed to use the api end point: %s: %s",
			userLevel, r.URL.Path, r.Method,
		)
		return errors.New(msg), 403, nil
	}

	return nil, 200, claims
}

func hasPermission(userLevel int, requiredLevel int) bool {
	return userLevel >= requiredLevel
}
