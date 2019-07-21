package local_account

import (
	"github.com/agiledragon/transfer-money-go/domain/model/base"
	"github.com/agiledragon/transfer-money-go/domain/model/common_role"
)

type LocalMoneyDest struct {
	base.Role
	accountInfo *common_role.AccountInfo
	balance     *balance
	phone       *phone
}

func (this *LocalMoneyDest) TransferMoneyFrom(fromId string, amount uint) {
	this.balance.increase(amount)
	this.phone.sendTransferFromMsg(fromId, amount)
}

func (this *LocalMoneyDest) GetAccountId() string {
	return this.accountInfo.Id()
}
