package models

type People interface {
	ExistUserByName() (bool, error)
	AddUser() error
	AuthorizedByUsernameAndPassword() (bool, error)
	GetRecords() ([]RecordResult, error)
}
