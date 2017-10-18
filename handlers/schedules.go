package handlers

import (
	"log"

	r "github.com/GoRethink/gorethink"
	"github.com/gombio/batory/db"
	"github.com/gombio/batory/ws"
)

// SchedulesNew creates new Schedule
func SchedulesCreate(s *ws.Session, data interface{}) {
	d := data.(map[string]interface{})
	schedule := db.NewSchedule(d["project"].(string), d["person"].(string), d["start"].(string), d["end"].(string))
	err := r.Table("schedules").Insert(schedule).Exec(s.Rethink.Session)
	if nil != err {
		log.Print(err.Error())
		s.Send <- ws.Message{Type: "error", Data: err.Error()} //TODO: Handle errors in clients

		return
	}

	log.Print(s.User.ID, ": schedules.create.success: ", s.User.ID)
	s.Send <- ws.Message{Type: "schedules.create.success", Data: schedule}
	return
}

// SchedulesList initiates listing of schedules
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
