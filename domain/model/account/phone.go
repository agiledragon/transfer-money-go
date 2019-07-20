package account

import (
	"fmt"
	"github.com/agiledragon/transfer-money-go/domain/model/base"
)

type phone struct {
	base.ValueObject
	accountInfo *accountInfo
	phoneNumber string
}

func newPhone(info *accountInfo, phoneNumber string) *phone {
	return &phone{accountInfo: info, phoneNumber: phoneNumber}
}

func (this *phone) sendWithdrawMsg(amount uint) {
	fmt.Printf("phone(%v): %v yuan has been withdrawed from accountId(%v)!\n",
		this.phoneNumber, amount, this.accountInfo.Id())
}
