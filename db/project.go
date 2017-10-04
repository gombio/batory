package db

type Project struct {
	Account string `gorethink:"account"`
	ID      string `gorethink:"id,omitempty"`
	Name    string `gorethink:"name"`
}

func NewProject(account string, id string, name string) Project {
	return Project{
		Account: account,
		ID:      id,
		Name:    name,
	}
}
