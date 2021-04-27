package repository

type DAO interface {
	Commit() error
	Rollback() error
	Begin() error
}
