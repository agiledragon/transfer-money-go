package service

import "github.com/agiledragon/transfer-money-go/domain/model/account"

type WithdrawMoneyService struct {
	repo account.AccountRepo
}

func NewWithdrawMoneyService() *WithdrawMoneyService {
	s := &WithdrawMoneyService{repo: account.GetAccountRepo()}
	return s
}

func (this *WithdrawMoneyService) Exec(accountId string, amount uint) {
	account := this.repo.Get(accountId)
	account.MoneyCollector.Withdraw(amount)
}
