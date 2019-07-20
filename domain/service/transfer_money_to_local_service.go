package service

import "github.com/agiledragon/transfer-money-go/domain/model/account"

type TransferMoneyToLocalService struct {
	repo account.AccountRepo
}

func NewTransferMoneyToLocalService() *TransferMoneyToLocalService {
	s := &TransferMoneyToLocalService{repo: account.GetAccountRepo()}
	return s
}

func (this *TransferMoneyToLocalService) Exec(fromId, toId string, amount uint) {
	fromAccount := this.repo.Get(fromId)
	toAccount := this.repo.Get(toId)
	fromAccount.MoneySrc.TransferMoneyTo(toAccount.MoneyDest, amount)
}
