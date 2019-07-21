package service

import (
	"github.com/agiledragon/transfer-money-go/domain/model/local_account"
	"github.com/agiledragon/transfer-money-go/domain/model/remote_account"
)

type TransferMoneyToRemoteService struct {
	repo local_account.LocalAccountRepo
}

func NewTransferMoneyToRemoteService() *TransferMoneyToRemoteService {
	s := &TransferMoneyToRemoteService{repo: local_account.GetLocalAccountRepo()}
	return s
}

func (this *TransferMoneyToRemoteService) Exec(fromId, toId string, amount uint) {
	fromAccount := this.repo.Get(fromId)
	toAccount := remote_account.New(toId)
	fromAccount.MoneySrc.TransferMoneyTo(toAccount.MoneyDest, amount)
}
