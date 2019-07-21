package remote_account

import (
	"github.com/agiledragon/transfer-money-go/domain/model/base"
	"github.com/agiledragon/transfer-money-go/domain/model/common_role"
)

type RemoteMoneySrc struct {
	base.Role
	accountInfo *common_role.AccountInfo
}

func (this *RemoteMoneySrc) TransferMoneyTo(dest common_role.MoneyDest, amount uint) {
	dest.TransferMoneyFrom(this.accountInfo.Id(), amount)
	ok := Response{false}
	sendProtocolResponse(ok)
}
