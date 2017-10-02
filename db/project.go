package db

type Project struct {
	Project string `gorethink:"project"`
	ID      string `gorethink:"id,omitempty"`
	Name    string `gorethink:"name"`
}

func NewProject(project string, id string, name string) Project {
	return Project{
		Project: project,
		ID:      id,
		Name:    name,
	}
}
