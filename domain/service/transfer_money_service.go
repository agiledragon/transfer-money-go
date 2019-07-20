package service

import "github.com/agiledragon/transfer-money-go/domain/model/account"

type TransferMoneyService struct {
	repo account.AccountRepo
}

func NewTransferMoneyService() *TransferMoneyService {
	s := &TransferMoneyService{repo: account.GetAccountRepo()}
	return s
}

func (this *TransferMoneyService) GetAmount(accountId string) uint {
	account := this.repo.Get(accountId)
	return account.MoneySrc.GetAmount()
}
