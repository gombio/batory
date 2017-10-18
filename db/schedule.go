package db

type Schedule struct {
	Project string `gorethink:"project" json:"project"`
	ID      string `gorethink:"id,omitempty" json:"id"`
	Person  string `gorethink:"person" json:"person"`
	Start   string `gorethink:"start" json:"start"`
	End     string `gorethink:"end" json:"end"`
}

func NewSchedule(project string, person string, start string, end string) Schedule {
	return Schedule{
		Project: project,
		Person:  person,
		Start:   start,
		End:     end,
	}
}
