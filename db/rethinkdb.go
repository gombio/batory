package db

import (
	r "github.com/GoRethink/gorethink"
)

type Rethinkdb struct {
	DbName  string
	Session *r.Session
}

func NewRethinkdb(DbName string, s *r.Session) *Rethinkdb {
	return &Rethinkdb{
		DbName:  DbName,
		Session: s,
	}
}

func (r *Rethinkdb) IsCreated(v r.ChangeResponse) bool {
	if nil != v.NewValue && nil == v.OldValue {
		return true
	}

	return false
}

func (r *Rethinkdb) IsUpdated(v r.ChangeResponse) bool {
	if nil != v.NewValue && nil != v.OldValue {
		return true
	}

	return false
}

func (r *Rethinkdb) IsDeleted(v r.ChangeResponse) bool {
	if nil == v.NewValue && nil != v.OldValue {
		return true
	}

	return false
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
		"users": {
			NewUser("gombio", "john@example.com", "John Doe", "john@example.com", "+11231231212", "test123", []string{}),
			NewUser("gombio", "juda@example.com", "Juda Smith", "juda@example.com", "+11231231212", "test123", []string{"project_batory_view"}),
			NewUser("acme", "joseph@example.com", "Joseph Fisher", "joseph@example.com", "+11231231212", "test123", []string{}),
		},
		"schedules": {
			//John, Youdash, Monday-Friday, 5:30-8:00
			NewSchedule("youdash", "john@example.com", "2017-10-02 05:30:00", "2017-10-02 08:00:00"),
			NewSchedule("youdash", "john@example.com", "2017-10-03 05:30:00", "2017-10-03 08:00:00"),
			NewSchedule("youdash", "john@example.com", "2017-10-04 05:30:00", "2017-10-04 08:00:00"),
			NewSchedule("youdash", "john@example.com", "2017-10-05 05:30:00", "2017-10-05 08:00:00"),
			NewSchedule("youdash", "john@example.com", "2017-10-06 05:30:00", "2017-10-06 08:00:00"),
		},
	}

	//Drop DB
	if dropExisting {
		r.DBDrop(s.DbName).Run(s.Session)
	}
	//Craete DB
	r.DBCreate(s.DbName).Run(s.Session)

	//Create tables with records
	for table, rows := range structure {
		r.DB(s.DbName).TableCreate(table).Run(s.Session)
		for _, row := range rows {
			r.DB(s.DbName).Table(table).Insert(row).Exec(s.Session)
		}
	}
}
