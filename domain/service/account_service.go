package service

import "github.com/agiledragon/transfer-money-go/domain/model/account"

type AccountService struct {
	repo account.AccountRepo
}

func NewAccountService() *AccountService {
	s := &AccountService{repo: account.GetAccountRepo()}
	return s
}

func (this *AccountService) CreateAccount(accountId, phoneNumber string, amount uint) *account.Account {
	a := account.NewAccount(accountId, phoneNumber, amount)
	this.repo.Add(a)
	return a
}

func (this *AccountService) DestoryAccount(accountId string) {
	this.repo.Remove(accountId)
}
