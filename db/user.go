package db

type User struct {
	Account     string   `gorethink:"account" json:"account"`
	ID          string   `gorethink:"id,omitempty" json:"id"`
	Name        string   `gorethink:"name" json:"name"`
	Email       string   `gorethink:"email" json:"email"`
	Phone       string   `gorethink:"phone" json:"phone"`
	Password    string   `gorethink:"password" json:"-"` //XXX: Ommit sending back the password
	Permissions []string `gorethink:"permissions" json:"permissions"`
}

func NewUser(account string, id string, name string, email string, phone string, password string, permissions []string) User {
	return User{
		Account:     account,
		ID:          id,
		Name:        name,
		Email:       email,
		Phone:       phone,
		Password:    password,
		Permissions: permissions,
	}
}
