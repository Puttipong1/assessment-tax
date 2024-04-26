package db

import (
	"github.com/Puttipong1/assessment-tax/common"
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/model"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/shopspring/decimal"
)

type deduction struct {
	deduction string  `postgres:"type"`
	amount    float64 `postgres:"amount"`
}

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

func (db *DB) GetDeductions() (model.Deduction, error) {
	log := config.Logger()
	d := model.Deduction{}
	query := "SELECT type, amount FROM deductions WHERE type IN ($1, $2, $3)"
	rows, err := db.DB.Query(query, common.PersonalDeductionsType, common.KReceiptDeductions, common.DonationsDeductionsType)
	if err != nil {
		log.Error().Err(err).Msg("Can't get deductions from database")
		return d, err
	}
	var deductions []deduction
	for rows.Next() {
		var deduction deduction
		err := rows.Scan(&deduction.deduction, &deduction.amount)
		if err != nil {
			log.Error().Err(err).Msg("")
			return d, err
		}
		deductions = append(deductions, deduction)
	}
	for _, deduction := range deductions {
		if deduction.deduction == common.PersonalDeductionsType {
			d.Personal = decimal.NewFromFloat(deduction.amount)
		}
		if deduction.deduction == common.KReceiptDeductionsType {
			d.KReceipt = decimal.NewFromFloat(deduction.amount)
		}
		if deduction.deduction == common.DonationsDeductionsType {
			d.Donation = decimal.NewFromFloat(deduction.amount)
		}

	}
	log.Info().Msgf("DB: Personal Deduction: %s, K-Receipt Deduction: %s", d.Personal, d.KReceipt)
	return d, nil
}
