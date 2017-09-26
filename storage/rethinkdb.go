package storage

import (
	"fmt"
	"github.com/gombio/batory/config"
	"github.com/gombio/batory/organization"
	"github.com/sirupsen/logrus"
	rethink "gopkg.in/gorethink/gorethink.v3"
)

type rethinkDb struct {
	session *rethink.Session
}

func newRethinkDb() (*rethinkDb, error) {
	options := rethink.ConnectOpts{
		Address:  config.AppConfig.Database.Address,
		Database: config.AppConfig.Database.DbName,
	}
	ses, err := rethink.Connect(options)
	if err != nil {
		return nil, err
	}
	return &rethinkDb{session: ses}, nil
}

func (r *rethinkDb) AddAccount(account organization.Account) (string, error) {
	res, err := rethink.Table("accounts").Insert(account).RunWrite(r.session)
	if err != nil {
		logrus.Error(err)
		return "", err
	}
	fmt.Printf("inserted %v \n", res.GeneratedKeys[0])
	return res.GeneratedKeys[0], nil

}

func (r *rethinkDb) ListProjects() {
	res, err := rethink.Table("projects").Run(r.session)
	if err != nil {
		logrus.Error(err)
	}
	defer res.Close()
	var hero []interface{}
	res.All(&hero)
	fmt.Printf("count %v \n", hero)
}

func (r *rethinkDb) GetById(id string) {
	res, err := rethink.Table("projects").Get(id).Run(r.session)
	if err != nil {
		logrus.Error(err)
	}
	defer res.Close()
	var hero interface{}
	res.One(&hero)
	fmt.Printf("count %v \n", hero)
}

func (r *rethinkDb) CreateDb() {
	res, err := rethink.DBCreate(config.AppConfig.Database.DbName).Run(r.session)
	if err != nil {
		logrus.Error(err)
	}
	defer res.Close()
}
