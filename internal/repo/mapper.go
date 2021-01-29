package repo

import "github.com/tupyy/stock/models"

func redisToStockValues(count int64, stockLabel string, values []float64) models.StockValues {
	return models.StockValues{
		Count:  count,
		Label:  stockLabel,
		Values: values,
	}
}
