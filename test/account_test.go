package test

import (
	"github.com/agiledragon/transfer-money-go/app/service"
	"github.com/agiledragon/transfer-money-go/domain/model/local_account"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type FakeAccountRepo struct {
	accounts map[string]*local_account.LocalAccount
}

func (this *FakeAccountRepo) Add(account *local_account.LocalAccount) {
	this.accounts[account.Id()] = account
}

func (this *FakeAccountRepo) Get(accountId string) *local_account.LocalAccount {
	return this.accounts[accountId]
}

//simulate dbs operation
func (this *FakeAccountRepo) Update(account *local_account.LocalAccount) {
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
	repo := &FakeAccountRepo{make(map[string]*local_account.LocalAccount)}
	local_account.SetLocalAccountRepo(repo)
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
		Convey("transfer money to remote", func() {
			const AMOUNT = 1000
			api.TransferMoneyToRemote(jimAccountId, lucyAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, jimInitialAmount-AMOUNT)
		})
		Convey("transfer money from remote", func() {
			const AMOUNT = 1000
			api.TransferMoneyFromRemote(lucyAccountId, jimAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, jimInitialAmount+AMOUNT)
		})
	})
}
