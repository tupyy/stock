package repo

import "github.com/tupyy/stock/models"

type StockRepo interface {

	// Add adds a value to the underling repo
	Add(stockValue models.StockValue)

	// GetStock returns the last saved stock value of a company
	GetStock(label string) (models.StockValue, error)

	// GetStocks returns all saved stock values of a company
	GetStocks(label string) (models.StockValues, error)

	// GetLatestStocks returns the latest values of each company
	GetLatestStocks() (count int64, values []*models.StockValue)
}
