package local_account

import (
	"github.com/agiledragon/transfer-money-go/domain/model/base"
	"github.com/agiledragon/transfer-money-go/domain/model/common_role"
)

type LocalAccount struct {
	base.AggregateRoot
	accountInfo    common_role.AccountInfo
	balance        balance
	phone          phone
	MoneyCollector MoneyCollector
	MoneySrc       common_role.MoneySrc
	MoneyDest      common_role.MoneyDest
}

func New(accountId, phoneNumber string, amount uint) *LocalAccount {
	account := &LocalAccount{
		AggregateRoot: base.NewAggregateRoot(accountId),
		accountInfo: common_role.AccountInfo{
			Entity: base.NewEntity(accountId),
		},
		balance: balance{
			ValueObject: base.ValueObject{},
			amount:      amount,
		},
	}
	account.phone = phone{
		ValueObject: base.ValueObject{},
		accountInfo: &account.accountInfo,
		phoneNumber: phoneNumber,
	}
	account.MoneyCollector = MoneyCollector{
		Role:    base.Role{},
		balance: &account.balance,
		phone:   &account.phone,
	}
	account.MoneySrc = &LocalMoneySrc{
		Role:        base.Role{},
		accountInfo: &account.accountInfo,
		balance:     &account.balance,
		phone:       &account.phone,
	}
	account.MoneyDest = &LocalMoneyDest{
		Role:        base.Role{},
		accountInfo: &account.accountInfo,
		balance:     &account.balance,
		phone:       &account.phone,
	}
	return account
}
