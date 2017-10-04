package db

type Schedule struct {
	Project string `gorethink:"project"`
	ID      string `gorethink:"id,omitempty"`
	Person  string `gorethink:"person"`
	Start   string `gorethink:"start"`
	End     string `gorethink:"end"`
}

func NewSchedule(project string, person string, start string, end string) Schedule {
	return Schedule{
		Project: project,
		Person:  person,
		Start:   start,
		End:     end,
	}
}
