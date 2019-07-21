package service

import (
	"github.com/agiledragon/transfer-money-go/domain/model/local_account"
	"github.com/agiledragon/transfer-money-go/domain/model/remote_account"
)

type TransferMoneyFromRemoteService struct {
	repo local_account.LocalAccountRepo
}

func NewTransferMoneyFromRemoteService() *TransferMoneyFromRemoteService {
	s := &TransferMoneyFromRemoteService{repo: local_account.GetLocalAccountRepo()}
	return s
}

func (this *TransferMoneyFromRemoteService) Exec(fromId, toId string, amount uint) {
	fromAccount := remote_account.New(fromId)
	toAccount := this.repo.Get(toId)
	fromAccount.MoneySrc.TransferMoneyTo(toAccount.MoneyDest, amount)
}
