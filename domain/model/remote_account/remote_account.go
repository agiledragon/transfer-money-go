package remote_account

import (
	"github.com/agiledragon/transfer-money-go/domain/model/base"
	"github.com/agiledragon/transfer-money-go/domain/model/common_role"
)

type RemoteAccount struct {
	base.AggregateRoot
	accountInfo common_role.AccountInfo
	MoneySrc    common_role.MoneySrc
	MoneyDest   common_role.MoneyDest
}

func New(accountId string) *RemoteAccount {
	account := &RemoteAccount{
		AggregateRoot: base.NewAggregateRoot(accountId),
		accountInfo: common_role.AccountInfo{
			Entity: base.NewEntity(accountId),
		},
	}
	account.MoneySrc = &RemoteMoneySrc{
		Role:        base.Role{},
		accountInfo: &account.accountInfo,
	}
	account.MoneyDest = &RemoteMoneyDest{
		Role:        base.Role{},
		accountInfo: &account.accountInfo,
	}
	return account
}
