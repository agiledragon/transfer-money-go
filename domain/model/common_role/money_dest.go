package common_role

type MoneyDest interface {
	TransferMoneyFrom(fromId string, amount uint)
	GetAccountId() string
}