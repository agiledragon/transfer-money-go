package local_account

import (
	"fmt"
	"github.com/agiledragon/transfer-money-go/domain/model/base"
	"github.com/agiledragon/transfer-money-go/domain/model/common_role"
)

type phone struct {
	base.ValueObject
	accountInfo *common_role.AccountInfo
	phoneNumber string
}

func (this *phone) sendWithdrawMsg(amount uint) {
	fmt.Printf("phone(%v): %v yuan has been withdrawed from accountId(%v)!\n",
		this.phoneNumber, amount, this.accountInfo.Id())
}

func (this *phone) sendTransferToMsg(toId string, amount uint) {
	fmt.Printf("phone(%v): accountId(%v) send %v yuan to accountId(%v)!\n",
		this.phoneNumber, this.accountInfo.Id(), amount, toId)
}

func (this *phone) sendTransferFromMsg(fromId string, amount uint) {
	fmt.Printf("phone(%v): accountId(%v) receive %v yuan from accountId(%v)!\n",
		this.phoneNumber, this.accountInfo.Id(), amount, fromId)
}
