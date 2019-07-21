package service

import (
	"github.com/agiledragon/transfer-money-go/domain/model/local_account"
)

type AccountService struct {
	repo local_account.LocalAccountRepo
}

func NewAccountService() *AccountService {
	s := &AccountService{repo:local_account.GetLocalAccountRepo()}
	return s
}

func (this *AccountService) CreateAccount(accountId, phoneNumber string, amount uint) *local_account.LocalAccount {
	a := local_account.New(accountId, phoneNumber, amount)
	this.repo.Add(a)
	return a
}

func (this *AccountService) DestoryAccount(accountId string) {
	this.repo.Remove(accountId)
}
