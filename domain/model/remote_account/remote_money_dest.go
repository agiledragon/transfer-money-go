package remote_account

import (
	"github.com/agiledragon/transfer-money-go/domain/model/base"
	"github.com/agiledragon/transfer-money-go/domain/model/common_role"
)

type RemoteMoneyDest struct {
	base.Role
	accountInfo *common_role.AccountInfo
}

func (this *RemoteMoneyDest) TransferMoneyFrom(fromId string, amount uint) {
	sendTransferToProtocolMsg(fromId, this.accountInfo.Id(), amount)
	resp := waitProtocolResp()
	if resp.isFailed {
		panic("transfer money to remote fail")
	}
}

func (this *RemoteMoneyDest) GetAccountId() string {
	return this.accountInfo.Id()
}
