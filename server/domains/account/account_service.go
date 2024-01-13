package account

import (
	"app/sql/db"
)

type IAccountService interface {
	GetAccount(id int) (db.Account, error)
	GetAccountByEmail(email string) (db.Account, error)
	CreateAccount(args db.CreateAccountParams) (db.Account, error)
}

type AccountService struct {
	repo IAccountRepo
}

func NewAccountService(repo IAccountRepo) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) CreateAccount(args db.CreateAccountParams) (db.Account, error) {
	return s.repo.Create(args)
}

func (s *AccountService) GetAccount(accountID int) (db.Account, error) {
	return s.repo.Get(int64(accountID))
}

func (s *AccountService) GetAccountByEmail(email string) (db.Account, error) {
	return s.repo.GetByEmail(email)
}
