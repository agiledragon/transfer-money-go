package account

import "github.com/agiledragon/transfer-money-go/domain/model/base"

type accountInfo struct {
	base.Entity
}

func newAccountInfo(id string) *accountInfo {
	return &accountInfo{base.NewEntity(id)}
}
