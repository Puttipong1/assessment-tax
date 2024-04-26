package config

import "os"

type DBConfig struct {
	user     string
	password string
	name     string
	url      string
	port     string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		user:     os.Getenv("DATABASE_USER"),
		password: os.Getenv("DATABASE_PASSWORD"),
		name:     os.Getenv("DATABASE_NAME"),
		url:      os.Getenv("DATABASE_URL"),
		port:     os.Getenv("DATABASE_PORT"),
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
func (dbConfig *DBConfig) Url() string {
	return dbConfig.url
}
func (dbConfig *DBConfig) Port() string {
	return dbConfig.port
}
