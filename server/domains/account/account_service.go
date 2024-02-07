package account

import (
	"app/sql/db"
	"strings"
)

type IAccountService interface {
	GetAccount(id int) (db.Account, error)
	GetAccountOrCreateIfNotExists(info UserInfoDto) (db.Account, error)
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
