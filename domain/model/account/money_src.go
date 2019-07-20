package account

import "github.com/agiledragon/transfer-money-go/domain/model/base"

type MoneySrc struct {
	base.Role
	accountInfo *accountInfo
	balance     *balance
	phone       *phone
}

func (this *MoneySrc) TransferMoneyTo(dest MoneyDest, amount uint) {
	if this.balance.get() < amount {
		panic("insufficient money!")
	}
	dest.transferMoneyFrom(this.accountInfo.Id(), amount)
	this.balance.decrease(amount)
	this.phone.sendTransferToMsg(dest.getAccountId(), amount)
}

func (this *MoneySrc) GetAmount() uint {
	return this.balance.get()
}
