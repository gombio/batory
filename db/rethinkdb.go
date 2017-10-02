package db

import (
	r "github.com/GoRethink/gorethink"
)

type Rethinkdb struct {
	dbName  string
	session *r.Session
}

func NewRethinkdb(dbName string, s *r.Session) *Rethinkdb {
	return &Rethinkdb{
		dbName:  dbName,
		session: s,
	}
}

func (s *Rethinkdb) Create(dropExisting bool) {
	structure := map[string][]interface{}{
		"accounts": {
			NewAccount("gombio", "Gomb.io"),
			NewAccount("acme", "ACME INC."),
		},
		"projects": {
			NewProject("gombio", "youdash", "Youdash"),
			NewProject("gombio", "batory", "Batory"),
			NewProject("acme", "example", "Example"),
		},
	}

	//Drop DB
	if dropExisting {
		r.DBDrop(s.dbName).Run(s.session)
	}
	//Craete DB
	r.DBCreate(s.dbName).Run(s.session)

	//Create tables with records
	for table, rows := range structure {
		r.DB(s.dbName).TableCreate(table).Run(s.session)
		for _, row := range rows {
			r.DB(s.dbName).Table(table).Insert(row).Exec(s.session)
		}
	}
}
