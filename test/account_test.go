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

	const JIM_PHONE_NUMBER = "19999999999"
	const JIM_INITIAL_AMOUNT = 10000
	const LUCY_PHONE_NUMBER = "18888888888"
	const LUCY_INITIAL_AMOUNT = 5000
	const AMOUNT = 1000

	Convey("TestAccount", t, func() {
		jimAccountId := api.CreateAccount(JIM_PHONE_NUMBER, JIM_INITIAL_AMOUNT)
		defer api.DestroyAccount(jimAccountId)
		lucyAccountId := api.CreateAccount(LUCY_PHONE_NUMBER, LUCY_INITIAL_AMOUNT)
		defer api.DestroyAccount(lucyAccountId)

		Convey("withdraw money", func() {
			api.WithdrawMoney(jimAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, JIM_INITIAL_AMOUNT-AMOUNT)
		})

		Convey("transfer money to local", func() {
			api.TransferMoneyToLocal(jimAccountId, lucyAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, JIM_INITIAL_AMOUNT-AMOUNT)
			So(api.GetAmount(lucyAccountId), ShouldEqual, LUCY_INITIAL_AMOUNT+AMOUNT)
		})
		Convey("transfer money to remote", func() {
			api.TransferMoneyToRemote(jimAccountId, lucyAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, JIM_INITIAL_AMOUNT-AMOUNT)
		})
		Convey("transfer money from remote", func() {
			api.TransferMoneyFromRemote(lucyAccountId, jimAccountId, AMOUNT)
			So(api.GetAmount(jimAccountId), ShouldEqual, JIM_INITIAL_AMOUNT+AMOUNT)
		})
	})
}
