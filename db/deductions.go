package db

import (
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/model/response"
)

func (db *DB) UpdateDeductions(deductionsType string, amount float64) error {
	log := config.Logger()
	query := "UPDATE deductions SET amount = $1 WHERE type = $2"
	res, err := db.DB.Exec(query, amount, deductionsType)
	if err != nil {
		log.Error().Err(err).Msg("Can't update deductions")
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("Can't update deductions")
		return err
	}
	if count == 0 {
		log.Error().Msgf("Row affects = %d", count)
		return &response.Error{}
	}
	return nil
}
