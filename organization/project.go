package organization

import (
	"github.com/satori/go.uuid"
)

type Project struct {
	id     uuid.UUID `gorethink:"id"`
	name   string    `gorethink:"name"`
	active bool
}

func NewProject(name string) *Project {
	uuid := uuid.NewV4()
	project := Project{id: uuid, name: name, active: true}
	return &project
}
