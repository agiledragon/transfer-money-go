package service

import "github.com/agiledragon/transfer-money-go/domain/model/local_account"

type TransferMoneyService struct {
	repo local_account.LocalAccountRepo
}

func NewTransferMoneyService() *TransferMoneyService {
	s := &TransferMoneyService{repo:local_account.GetLocalAccountRepo()}
	return s
}

func (this *TransferMoneyService) GetAmount(accountId string) uint {
	account := this.repo.Get(accountId)
	return account.MoneySrc.(*local_account.LocalMoneySrc).GetAmount()
}
