package account

import (
	"app/sql/db"
	"net/http"
	"strings"
)

type IAccountService interface {
	GetAccount(r *http.Request, id int) (db.Account, error)
	GetAccountOrCreateIfNotExists(info UserInfoDto) (db.Account, error)
	GetAccountByEmail(email string) (db.Account, error)
	CreateAccount(args db.CreateAccountParams) (db.Account, error)
}

type AccountService struct {
	repo     IAccountRepo
	authServ IAuthService
}

func NewAccountService(repo IAccountRepo, authServ IAuthService) *AccountService {
	return &AccountService{
		repo:     repo,
		authServ: authServ,
	}
}

func (s *AccountService) CreateAccount(args db.CreateAccountParams) (db.Account, error) {
	return s.repo.Create(args)
}

func (s *AccountService) GetAccount(r *http.Request, accountID int) (db.Account, error) {
	acc, err := s.repo.Get(int64(accountID))
	if err != nil {
		return db.Account{}, err
	}
	token, err := s.authServ.JwtFrom(r)
	if err != nil {
		return db.Account{}, err
	}
	userID := s.authServ.CurrentUserID(token)
	// do not disclose email of another user to the currently logged in user.
	if acc.Email != userID {
		acc.Email = "Undisclosed"
	}
	return acc, nil
}

func (s *AccountService) GetAccountByEmail(email string) (db.Account, error) {
	acc, err := s.repo.GetByEmail(email)
	if err != nil {
		return db.Account{}, err
	}
	return acc, nil
}

func (s *AccountService) GetAccountOrCreateIfNotExists(info UserInfoDto) (db.Account, error) {
	acc, err := s.GetAccountByEmail(info.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "no rows") {
			return db.Account{}, err
		}
		acc, err = s.CreateAccount(db.CreateAccountParams{
			Name:  info.Name,
			Level: 0,
			Email: info.Email,
		})
		if err != nil {
			return db.Account{}, err
		}
	}
	return acc, nil
}
