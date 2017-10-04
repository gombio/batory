package db

type User struct {
	Account     string   `gorethink:"account"`
	ID          string   `gorethink:"id,omitempty"`
	Name        string   `gorethink:"name"`
	Email       string   `gorethink:"email"`
	Phone       string   `gorethink:"phone"`
	Password    string   `gorethink:"password"`
	Permissions []string `gorethink:"permissions"`
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
