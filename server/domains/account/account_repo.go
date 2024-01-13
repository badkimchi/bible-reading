package account

import (
	"app/sql/db"
	"context"
)

type IAccountRepo interface {
	Get(id int64) (db.Account, error)
	GetByEmail(email string) (db.Account, error)
	Create(args db.CreateAccountParams) (db.Account, error)
}

type AccountRepo struct {
	db *db.Queries
}

func NewAccountRepo(db *db.Queries) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (r *AccountRepo) Get(accountID int64) (db.Account, error) {
	acc, err := r.db.GetAccount(context.Background(), accountID)
	if err != nil {
		return db.Account{}, err
	}
	return acc, nil
}

func (r *AccountRepo) GetByEmail(email string) (db.Account, error) {
	acc, err := r.db.GetAccountByEmail(context.Background(), email)
	if err != nil {
		return db.Account{}, err
	}
	return acc, nil
}

func (r *AccountRepo) Create(args db.CreateAccountParams) (db.Account, error) {
	acc, err := r.db.CreateAccount(context.Background(), args)
	if err != nil {
		return db.Account{}, err
	}
	return acc, nil
}
