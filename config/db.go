package config

import "os"

type DBConfig struct {
	user     string
	password string
	name     string
	host     string
	port     string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		name:     os.Getenv("DB_NAME"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
	}
}

func (dbConfig *DBConfig) User() string {
	return dbConfig.user
}
func (dbConfig *DBConfig) Password() string {
	return dbConfig.password
}
func (dbConfig *DBConfig) Name() string {
	return dbConfig.name
}
func (dbConfig *DBConfig) Host() string {
	return dbConfig.host
}
func (dbConfig *DBConfig) Port() string {
	return dbConfig.port
}
