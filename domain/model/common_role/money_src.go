package common_role

type MoneySrc interface {
	TransferMoneyTo(dest MoneyDest, amount uint)
}
