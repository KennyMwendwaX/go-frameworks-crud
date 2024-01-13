package config

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

var DBConfig = DatabaseConfig{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "173890",
	Name:     "go-restapi-crud",
}
