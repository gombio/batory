package config

type webConfig struct {
	address string
	port    string
}
type dbConfig struct {
	Provider string
	Address  string
	Port     string
	DbName   string
}

var (
	AppConfig struct {
		Database dbConfig
		Web      webConfig
	}
)
