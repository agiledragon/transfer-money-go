package local_account

var localAccountRepo LocalAccountRepo

type LocalAccountRepo interface {
	Add(account *LocalAccount)
	Get(accountId string) *LocalAccount
	Update(account *LocalAccount)
	Remove(accountId string)
}

func SetLocalAccountRepo(repo LocalAccountRepo) {
	localAccountRepo = repo
}

func GetLocalAccountRepo() LocalAccountRepo {
	return localAccountRepo
}
