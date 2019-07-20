package account

import "github.com/agiledragon/transfer-money-go/domain/model/base"

type MoneyDest struct {
	base.Role
	accountInfo *accountInfo
	balance     *balance
	phone       *phone
}

func (this *MoneyDest) transferMoneyFrom(fromId string, amount uint) {
	this.balance.increase(amount)
	this.phone.sendTransferFromMsg(fromId, amount)
}

func (this *MoneyDest) getAccountId() string {
	return this.accountInfo.Id()
}
