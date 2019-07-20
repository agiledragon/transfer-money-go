package infra

import "github.com/agiledragon/transfer-money-go/domain/model/account"

type AccountRepoImpl struct {
}

func (this *AccountRepoImpl) Add(account *account.Account) {

}

func (this *AccountRepoImpl) Get(accountId string) *account.Account {
	return nil
}

func (this *AccountRepoImpl) Update(account *account.Account) {

}

func (this *AccountRepoImpl) Remove(accountId string) {

}
