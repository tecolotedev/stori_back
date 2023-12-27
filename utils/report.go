package utils

import (
	"strconv"

	"github.com/tecolotedev/stori_back/db/sqlc_code"
	"github.com/tecolotedev/stori_back/email"
)

type AccCreditDebit struct {
	credit []float64
	debit  []float64
}

func MakeSummary(transfers []sqlc_code.Transfer) []email.TransferSummary {

	helper := make(map[string]AccCreditDebit)
	var transferSummaries []email.TransferSummary

	for _, transfer := range transfers {
		year, month, _ := transfer.CreatedAt.Time.Date()

		yearStr := strconv.Itoa(year)
		monthStr := "0" + strconv.Itoa(int(month))

		monthStr = monthStr[len(monthStr)-2:] //take last 2 numbers

		date := monthStr + "/" + yearStr

		val, ok := helper[date]

		if !ok {
			val = AccCreditDebit{}
		}

		if transfer.Amount > 0 {
			val.credit = append(val.credit, transfer.Amount)
		} else {
			val.debit = append(val.debit, transfer.Amount)
		}

		helper[date] = val

	}

	for key, value := range helper {

		sumCredit := 0.0
		for _, v := range value.credit {
			sumCredit += v
		}

		sumDebit := 0.0
		for _, v := range value.debit {
			sumDebit += v
		}

		t1 := email.TransferSummary{
			Month:   key,
			Type:    "credit",
			Average: float64(sumCredit) / float64(len(value.credit)),
		}

		t2 := email.TransferSummary{
			Month:   key,
			Type:    "debit",
			Average: float64(sumDebit) / float64(len(value.debit)),
		}

		transferSummaries = append(transferSummaries, t1, t2)
	}

	return transferSummaries
}
