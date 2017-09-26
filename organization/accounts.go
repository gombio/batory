package organization

type Accounts struct {
	Accounts map[string]Account `gorethink:"name"`
}
type Account struct {
	Name     string             `gorethink:"name"`
	Owner    Owner              `gorethink:"owner"`
	Users    map[string]User    `gorethink:"user"`
	People   map[string]People  `gorethink:"people"`
	Projects map[string]Project `gorethink:"projects"`
}

type Owner struct {
	Name  string `gorethink:"name"`
	Email string `gorethink:"email"`
	Phone string `gorethink:"phone"`
}
type User struct {
	Name     string `gorethink:"name"`
	Email    string `gorethink:"email"`
	Phone    string `gorethink:"phone"`
	Password string `gorethink:"password"`
}

type People struct {
	Name  string `gorethink:"name"`
	Email string `gorethink:"email"`
	Phone string `gorethink:"phone"`
}

type Project struct {
	Name           string     `gorethink:"name"`
	OnCallSchedule []Schedule `gorethink:"onCallSchedule",yaml:"onCallSchedule"`
}

type Schedule struct {
	Person string `gorethink:"person"`
	Start  string `gorethink:"start"`
	End    string `gorethink:"end"`
}
