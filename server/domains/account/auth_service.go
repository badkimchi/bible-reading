package account

import (
	"github.com/go-chi/jwtauth"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IAuthService interface {
	setAuthTokenDuration(duration time.Duration)
	authTokenExpireTime() time.Time
	refreshTokenExpireTime() time.Time
	getJwt(accountID string, level int) Jwt
	exchangeRefreshToken(tokenString string) (bool, string, string)
	CurrentUserID(r *http.Request) string
}

type AuthService struct {
	tokenAuth            *jwtauth.JWTAuth
	accountServ          IAccountService
	authTokenDuration    time.Duration
	refreshTokenDuration time.Duration
}

func NewAuthService(
	tAuth *jwtauth.JWTAuth,
	accountServ IAccountService,
) *AuthService {
	return &AuthService{
		tokenAuth:            tAuth,
		accountServ:          accountServ,
		authTokenDuration:    time.Hour * 12,
		refreshTokenDuration: time.Hour * 13,
	}
}

func (s *AuthService) setAuthTokenDuration(duration time.Duration) {
	s.authTokenDuration = duration
}

func (s *AuthService) authTokenExpireTime() time.Time {
	return time.Now().Add(s.authTokenDuration)
}
func (s *AuthService) refreshTokenExpireTime() time.Time {
	return time.Now().Add(s.refreshTokenDuration)
}

func (s *AuthService) getJwt(userID string, level int) Jwt {
	authToken, expire := s.authToken(userID, level)
	refToken, refExpire := s.getRefreshToken(userID, level)
	rToken := RefreshToken{Token: refToken, Expiration: refExpire}
	return Jwt{Token: authToken, Expiration: expire, RefreshToken: rToken}
}

func (s *AuthService) CurrentUserID(r *http.Request) string {
	tokenStr := r.Header.Get("Authorization")
	token, found := strings.CutPrefix(tokenStr, "Bearer ")
	if !found {
		return "BAD_AUTH_TOKEN"
	}
	jwt, err := s.tokenAuth.Decode(token)
	if err != nil {
		return "BAD_AUTH_TOKEN-2"
	}
	claims := jwt.PrivateClaims()
	userID := claims["user_id"].(string)
	return userID
}

// LoginInfo id is embedded in
func (s *AuthService) authToken(userID string, level int) (string, string) {
	aTokenClaims := map[string]interface{}{
		"user_id":    userID,
		"token_type": "auth",
		"level":      strconv.Itoa(level),
	}
	jwtauth.SetExpiry(aTokenClaims, s.authTokenExpireTime())
	_, authToken, _ := s.tokenAuth.Encode(aTokenClaims)
	return authToken, s.authTokenExpireTime().String()
}

func (s *AuthService) getRefreshToken(userID string, level int) (string, string) {
	rTokenClaims := map[string]interface{}{
		"user_id":    userID,
		"token_type": "refresh",
		"level":      strconv.Itoa(level),
	}
	jwtauth.SetExpiry(rTokenClaims, s.refreshTokenExpireTime())
	_, refreshToken, _ := s.tokenAuth.Encode(rTokenClaims)
	return refreshToken, s.refreshTokenExpireTime().String()
}

func (s *AuthService) exchangeRefreshToken(tokenString string) (bool, string, string) {
	token, err := s.tokenAuth.Decode(tokenString)
	if err != nil {
		return false, err.Error(), ""
	}
	claims := token.PrivateClaims()
	if claims["token_type"] != "refresh" {
		return false, "This is not s refresh token", ""
	}
	levelStr := claims["level"].(string)
	level, _ := strconv.Atoi(levelStr)
	rToken, expirationTime := s.getRefreshToken(claims["user_id"].(string), level)
	return true, rToken, expirationTime
}
