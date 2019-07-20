package test

import (
	"github.com/agiledragon/transfer-money-go/app/service"
	"github.com/agiledragon/transfer-money-go/domain/model/account"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type FakeAccountRepo struct {
	accounts map[string]*account.Account
}

func (this *FakeAccountRepo) Add(account *account.Account) {
	this.accounts[account.Id()] = account
}

func (this *FakeAccountRepo) Get(accountId string) *account.Account {
	return this.accounts[accountId]
}

//simulate dbs operation
func (this *FakeAccountRepo) Update(account *account.Account) {
	for k := range this.accounts {
		if k == account.Id() {
			delete(this.accounts, k)
		}
	}
	this.accounts[account.Id()] = account
}

func (this *FakeAccountRepo) Remove(accountId string) {
	delete(this.accounts, accountId)
}

func TestAccount(t *testing.T) {
	repo := &FakeAccountRepo{make(map[string]*account.Account)}
	account.SetAccountRepo(repo)
	api := service.NewAccountApi()

	Convey("TestAccount", t, func() {
		jimPhoneNumber := "19999999999"
		jimInitialAmount := uint(10000)
		jimAccountId := api.CreateAccount(jimPhoneNumber, jimInitialAmount)
		defer api.DestroyAccount(jimAccountId)

		lucyPhoneNumber := "18888888888"
		lucyInitialAmount := uint(5000)
		lucyAccountId := api.CreateAccount(lucyPhoneNumber, lucyInitialAmount)
		defer api.DestroyAccount(lucyAccountId)

		Convey("withdraw money", func() {
			const AMOUNT = 1000
			api.WithdrawMoney(jimAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, jimInitialAmount-AMOUNT)
		})

		Convey("transfer money to local", func() {
			const AMOUNT = 1000
			api.TransferMoneyToLocal(jimAccountId, lucyAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, jimInitialAmount-AMOUNT)
			So(api.GetAmount(lucyAccountId), ShouldEqual, lucyInitialAmount+AMOUNT)
		})
	})
}
