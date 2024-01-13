package account

import (
	"app/util/resp"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type AccountController struct {
	serv     IAccountService
	authServ IAuthService
}

func NewAccountController(
	serv IAccountService,
	authServ IAuthService,
) AccountController {
	return AccountController{
		serv:     serv,
		authServ: authServ,
	}
}

func (c *AccountController) GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if len(idStr) == 0 {
		resp.Bad(w, r, errors.New("id must be set"))
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}
	acc, err := c.serv.GetAccount(id)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}
	userID := c.authServ.CurrentUserID(r)
	if acc.Email != userID {
		acc.Email = "Undisclosed"
	}
	resp.Data(w, r, acc)
}
