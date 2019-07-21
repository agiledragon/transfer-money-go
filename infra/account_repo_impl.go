package infra

import (
	"github.com/agiledragon/transfer-money-go/domain/model/local_account"
)

type AccountRepoImpl struct {
}

func (this *AccountRepoImpl) Add(account *local_account.LocalAccount) {

}

func (this *AccountRepoImpl) Get(accountId string) *local_account.LocalAccount {
	return nil
}

func (this *AccountRepoImpl) Update(account *local_account.LocalAccount) {

}

func (this *AccountRepoImpl) Remove(accountId string) {

}
