package service

import "github.com/agiledragon/transfer-money-go/domain/model/local_account"

type WithdrawMoneyService struct {
	repo local_account.LocalAccountRepo
}

func NewWithdrawMoneyService() *WithdrawMoneyService {
	s := &WithdrawMoneyService{repo:local_account.GetLocalAccountRepo()}
	return s
}

func (this *WithdrawMoneyService) Exec(accountId string, amount uint) {
	account := this.repo.Get(accountId)
	account.MoneyCollector.Withdraw(amount)
}
