package service

import "github.com/agiledragon/transfer-money-go/domain/model/local_account"

type TransferMoneyToLocalService struct {
	repo local_account.LocalAccountRepo
}

func NewTransferMoneyToLocalService() *TransferMoneyToLocalService {
	s := &TransferMoneyToLocalService{repo: local_account.GetLocalAccountRepo()}
	return s
}

func (this *TransferMoneyToLocalService) Exec(fromId, toId string, amount uint) {
	fromAccount := this.repo.Get(fromId)
	toAccount := this.repo.Get(toId)
	fromAccount.MoneySrc.TransferMoneyTo(toAccount.MoneyDest, amount)
}
