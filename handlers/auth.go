package handlers

import (
	"log"

	r "github.com/GoRethink/gorethink"
	"github.com/gombio/batory/db"
	"github.com/gombio/batory/ws"
	"github.com/mitchellh/mapstructure"
)

//Auth401 returns 401 equivalent (authentication required) back to the session
func Auth401(s *ws.Session, data interface{}) {
	s.Send <- ws.Message{Type: "auth.required", Data: "Authentication required"}
	return
}

func AuthLogin(s *ws.Session, data interface{}) {
	creds := data.(map[string]interface{})
	cursor, err := r.Table("users").Filter(map[string]interface{}{
		"email":    creds["login"],
		"password": creds["password"],
	}).Run(s.Rethink.Session)
	if (nil != err) || cursor.IsNil() {
		s.Send <- ws.Message{Type: "auth.login.failure", Data: ""}
		return
	}

	var row interface{}
	for cursor.Next(&row) {
		var user db.User
		err := mapstructure.Decode(row, &user)
		if nil != err {
			s.Send <- ws.Message{Type: "error", Data: err.Error()}
			return
		}

		//Authenticate user/session
		s.IsAuthenticated = true
		s.User = user

		//NOTE: use s.ID here so that we can match SESSION with USER later, as from
		//now on s.Identifier() is going to returned USER.ID instead of SESSION.ID
		log.Print(s.ID, ": auth.login.success: ", s.User.ID)
		s.Send <- ws.Message{Type: "auth.login.success", Data: s.User}
		return
	}
}
