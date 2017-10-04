package handlers

import (
	"log"

	r "github.com/GoRethink/gorethink"
	"github.com/gombio/batory/ws"
)

func SchedulesList(s *ws.Session, data interface{}) {
	cursor, err := r.Table("schedules").Changes(r.ChangesOpts{IncludeInitial: true}).Run(s.Rethink.Session)
	if nil != err {
		log.Print(err.Error())
		s.Send <- ws.Message{Type: "error", Data: err.Error()} //TODO: Handle errors in clients

		return
	}

	result := make(chan r.ChangeResponse)
	go func() {
		var change r.ChangeResponse
		for cursor.Next(&change) {
			result <- change
		}
	}()

	go func() {
		for {
			select {
			// case <-stop:
			// 	cursor.Close()
			// 	return
			case change := <-result:
				if s.Rethink.IsCreated(change) {
					s.Send <- ws.Message{Type: "schedule.add", Data: change.NewValue}
				}
				// log.Printf("IsCreated: ", s.Rethink.IsCreated(change))
				// log.Printf("IsDeleted: ", s.Rethink.IsDeleted(change))
				// log.Printf("IsUpdated: ", s.Rethink.IsUpdated(change))
			}
		}
	}()
}
