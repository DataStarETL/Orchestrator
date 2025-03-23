package config

import "os"

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

type ServerConfig struct {
	Port string
}

func LoadConfig() (*DatabaseConfig, *ServerConfig) {
	dbConfig := DatabaseConfig{Host: os.Getenv("DB_HOST"), Port: os.Getenv("DB_PORT"), User: os.Getenv("DB_USERNAME"), Password: os.Getenv("DB_PASSWORD"), Dbname: os.Getenv("DB_DATABASE")}
	serverConfig := ServerConfig{Port: os.Getenv("REST_PORT")}
	return &dbConfig, &serverConfig
}
