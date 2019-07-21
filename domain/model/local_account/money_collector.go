package local_account

import "github.com/agiledragon/transfer-money-go/domain/model/base"

type MoneyCollector struct {
	base.Role
	balance *balance
	phone   *phone
}

func (this *MoneyCollector) Withdraw(amount uint) {
	if this.balance.get() < amount {
		panic("insufficient money!")
	}
	this.balance.decrease(amount)
	this.phone.sendWithdrawMsg(amount)
}

func (this *MoneyCollector) GetAmount() uint {
	return this.balance.get()
}
