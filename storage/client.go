package storage

import (
	"errors"
	"github.com/gombio/batory/config"
	"github.com/gombio/batory/organization"
)

type DataStorage interface {
	AccountCreator
}

type AccountCreator interface {
	AddAccount(account organization.Account) (string, error)
}

func NewDb() (DataStorage, error) {
	switch config.AppConfig.Database.Provider {
	case "rethinkDb":
		db, e := newRethinkDb()
		if e != nil {
			return nil, e
		}
		return db, nil
	}
	return nil, errors.New("err")
}
