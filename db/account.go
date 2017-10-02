package db

type Account struct {
	ID   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func NewAccount(id string, name string) Account {
	return Account{
		ID:   id,
		Name: name,
	}
}
