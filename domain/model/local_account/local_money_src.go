package local_account

import (
	"github.com/agiledragon/transfer-money-go/domain/model/base"
	"github.com/agiledragon/transfer-money-go/domain/model/common_role"
)

type LocalMoneySrc struct {
	base.Role
	accountInfo *common_role.AccountInfo
	balance     *balance
	phone       *phone
}

func (this *LocalMoneySrc) TransferMoneyTo(dest common_role.MoneyDest, amount uint) {
	if this.balance.get() < amount {
		panic("insufficient money!")
	}
	dest.TransferMoneyFrom(this.accountInfo.Id(), amount)
	this.balance.decrease(amount)
	this.phone.sendTransferToMsg(dest.GetAccountId(), amount)
}

func (this *LocalMoneySrc) GetAmount() uint {
	return this.balance.get()
}
