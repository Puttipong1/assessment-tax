package db

import (
	"database/sql"
	"fmt"

	"github.com/Puttipong1/assessment-tax/config"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

func Init(dbConfig *config.DBConfig) *DB {
	log := config.Logger()
	databaseSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Url(), dbConfig.Port(), dbConfig.User(), dbConfig.Password(), dbConfig.Name())
	db, err := sql.Open("postgres", databaseSource)
	if err != nil {
		log.Fatal().Msgf("Connect to database error %s", err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal().Msgf("Connect to database error %s", err.Error())
	}
	return &DB{DB: db}
}
