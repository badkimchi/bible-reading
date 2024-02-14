package account

import (
	"app/util/resp"
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type AccountController struct {
	serv IAccountService
}

func NewAccountController(
	serv IAccountService,
) AccountController {
	return AccountController{
		serv: serv,
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
	acc, err := c.serv.GetAccount(r, id)
	if err != nil {
		resp.Bad(w, r, err)
		return
	}
	resp.Data(w, r, acc)
}
