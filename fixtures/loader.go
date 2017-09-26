package fixtures

import (
	"fmt"
	"github.com/gombio/batory/organization"
	"github.com/sirupsen/logrus"
	//"os"
	"github.com/gombio/batory/storage"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadFixtures() {
	_, e := storage.NewDb()
	if e != nil {
		logrus.Error(e)
	}
	yamlFile, err := ioutil.ReadFile("../../fixtures/data.yaml")
	if err != nil {
		logrus.Panic("yamlFile.Get err   #%v ", err)
	}

	var accounts organization.Accounts
	err = yaml.Unmarshal(yamlFile, &accounts)
	if err != nil {
		logrus.Fatal("#%v ", err)
	}

	for _, account := range accounts.Accounts {
		db, _ := storage.NewDb()
		id, _ := db.AddAccount(account)
		logrus.Info(fmt.Sprintf("%s", id))
	}
}
