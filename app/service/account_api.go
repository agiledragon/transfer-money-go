package service

import "github.com/agiledragon/transfer-money-go/domain/service"

type AccountApi struct {
	accountService              *service.AccountService
	withdrawMoneyService        *service.WithdrawMoneyService
	transferMoneyToLocalService *service.TransferMoneyToLocalService
	transferMoneyService        *service.TransferMoneyService
	transferMoneyToRemoteService *service.TransferMoneyToRemoteService
}

func NewAccountApi() *AccountApi {
	api := &AccountApi{
		accountService:              service.NewAccountService(),
		withdrawMoneyService:        service.NewWithdrawMoneyService(),
		transferMoneyToLocalService: service.NewTransferMoneyToLocalService(),
		transferMoneyService:        service.NewTransferMoneyService(),
		transferMoneyToRemoteService:service.NewTransferMoneyToRemoteService(),
	}
	return api
}

func (this *AccountApi) generateAccountId(phoneNumber string) string {
	return phoneNumber[2:] + "888"
}

func (this *AccountApi) CreateAccount(phoneNumber string, initialAmount uint) string {
	id := this.generateAccountId(phoneNumber)
	this.accountService.CreateAccount(id, phoneNumber, initialAmount)
	return id
}

func (this *AccountApi) DestroyAccount(accountId string) {
	this.accountService.DestoryAccount(accountId)
}

func (this *AccountApi) WithdrawMoney(accountId string, amount uint) {
	this.withdrawMoneyService.Exec(accountId, amount)
}

func (this *AccountApi) GetAmount(accountId string) uint {
	return this.transferMoneyService.GetAmount(accountId)
}

func (this *AccountApi) TransferMoneyToLocal(fromId, toId string, amount uint) {
	this.transferMoneyToLocalService.Exec(fromId, toId, amount)
}

func (this *AccountApi) TransferMoneyToRemote(fromId, toId string, amount uint) {
	this.transferMoneyToRemoteService.Exec(fromId, toId, amount)
}