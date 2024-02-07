package account

import (
	"app/conf"
	"app/util/resp"
	"encoding/json"
	"errors"
	"github.com/go-chi/jwtauth"
	"net/http"
)

type IController interface {
}

type AuthController struct {
	config    *conf.Config
	tokenAuth *jwtauth.JWTAuth
	serv      IAuthService
	accServ   IAccountService
}

func NewAuthController(
	config *conf.Config,
	tokenAuth *jwtauth.JWTAuth,
	hAuth IAuthService,
	accServ IAccountService,
) AuthController {
	return AuthController{
		config:    config,
		tokenAuth: tokenAuth,
		serv:      hAuth,
		accServ:   accServ,
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tokenReq OAuthRequest
	err := decoder.Decode(&tokenReq)
	if err != nil {
		resp.Bad(w, r, errors.New("EOF: unable to parse token request"+err.Error()))
		return
	}
	if tokenReq.Token == "" {
		resp.Bad(w, r, errors.New("token must be passed in"))
		return
	}

	info, err := c.serv.GetUserInfo(tokenReq, c.config)
	acc, err := c.accServ.GetAccountOrCreateIfNotExists(info)
	info.Jwt = c.serv.getJwt(info.Email, int(acc.Level))
	resp.Data(w, r, info)
}

func (c *AuthController) RefreshWithRefreshToken(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req RefreshToken
	err := decoder.Decode(&req)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}

	_, refreshToken, refreshTokenExpiration := c.serv.exchangeRefreshToken(req.Token)
	rToken := RefreshToken{
		Token:      refreshToken,
		Expiration: refreshTokenExpiration,
	}
	resp.Data(w, r, rToken)
}
