package account

import "github.com/agiledragon/transfer-money-go/domain/model/base"

type Account struct {
	base.AggregateRoot
	accountInfo    accountInfo
	balance        balance
	phone          phone
	MoneyCollector MoneyCollector
	MoneySrc       MoneySrc
	MoneyDest      MoneyDest
}

func NewAccount(accountId, phoneNumber string, amount uint) *Account {
	account := &Account{
		AggregateRoot: base.NewAggregateRoot(accountId),
		accountInfo: accountInfo{
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
	account.MoneySrc = MoneySrc{
		Role:        base.Role{},
		accountInfo: &account.accountInfo,
		balance:     &account.balance,
		phone:       &account.phone,
	}
	account.MoneyDest = MoneyDest{
		Role:        base.Role{},
		accountInfo: &account.accountInfo,
		balance:     &account.balance,
		phone:       &account.phone,
	}
	return account
}
