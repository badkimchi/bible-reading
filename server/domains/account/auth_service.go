package account

import (
	"app/conf"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IAuthService interface {
	setAuthTokenDuration(duration time.Duration)
	authTokenExpireTime() time.Time
	refreshTokenExpireTime() time.Time
	createJwt(accountID string, level int) Jwt
	JwtFrom(r *http.Request) (jwt.Token, error)
	CurrentUserID(token jwt.Token) string
	exchangeRefreshToken(tokenString string) (bool, string, string)
	GetUserInfo(req OAuthRequest, config *conf.Config) (UserInfoDto, error)
}

type AuthService struct {
	tokenAuth            *jwtauth.JWTAuth
	authTokenDuration    time.Duration
	refreshTokenDuration time.Duration
}

func NewAuthService(
	tAuth *jwtauth.JWTAuth,
) *AuthService {
	return &AuthService{
		tokenAuth:            tAuth,
		authTokenDuration:    time.Hour * 12,
		refreshTokenDuration: time.Hour * 13,
	}
}

func (s *AuthService) GetUserInfo(req OAuthRequest, config *conf.Config) (UserInfoDto, error) {
	var Endpoint = oauth2.Endpoint{
		AuthURL:       "https://accounts.google.com/o/oauth2/auth",
		TokenURL:      "https://oauth2.googleapis.com/token",
		DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
		AuthStyle:     oauth2.AuthStyleInParams,
	}
	oauth := oauth2.Config{
		ClientID:     config.GoogleClientID,
		ClientSecret: config.GoogleClientSecret,
		Endpoint:     Endpoint,
		RedirectURL:  "postmessage",
		Scopes:       []string{"https://www.googleapis.com/auth/drive.metadata.readonly"},
	}
	token, err := oauth.Exchange(context.Background(), req.Token)
	if err != nil {
		return UserInfoDto{}, err
	}

	authReq, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v3/userinfo", bytes.NewBuffer([]byte("")))
	authReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	if err != nil {
		return UserInfoDto{}, err
	}
	client := &http.Client{Timeout: time.Millisecond * 1000}
	apiResp, err := client.Do(authReq)
	if err != nil {
		return UserInfoDto{}, err
	}

	decoder := json.NewDecoder(apiResp.Body)
	var info UserInfoDto
	err = decoder.Decode(&info)
	if err != nil {
		return UserInfoDto{}, err
	}
	return info, nil
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

func (s *AuthService) createJwt(userID string, level int) Jwt {
	authToken, expire := s.authToken(userID, level)
	refToken, refExpire := s.getRefreshToken(userID, level)
	rToken := RefreshToken{Token: refToken, Expiration: refExpire}
	return Jwt{Token: authToken, Expiration: expire, RefreshToken: rToken}
}

func (s *AuthService) CurrentUserID(token jwt.Token) string {
	claims := token.PrivateClaims()
	userID := claims["user_id"].(string)
	return userID
}

func (s *AuthService) JwtFrom(r *http.Request) (jwt.Token, error) {
	tokenStr := r.Header.Get("Authorization")
	token, found := strings.CutPrefix(tokenStr, "Bearer ")
	if !found {
		return nil, errors.New("no bearer token found")
	}
	jsonWebToken, err := s.tokenAuth.Decode(token)
	if err != nil {
		return nil, err
	}
	return jsonWebToken, nil
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
