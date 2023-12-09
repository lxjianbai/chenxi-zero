package utils

import "github.com/shopspring/decimal"

func DecimalCurrency(num float64) decimal.Decimal {
	return decimal.NewFromFloat(num).Round(15)
}

func DecimalDefault(num float64) decimal.Decimal {
	return decimal.NewFromFloat(num)
}

func DecimalWithRound(num float64, round uint32) decimal.Decimal {
	return decimal.NewFromFloat(num).Round(int32(round))
}
