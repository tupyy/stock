package repo

import (
	"fmt"
	"sync"

	"github.com/tupyy/stock/models"
)

type DeprecatedStockRepo struct {
	stocks map[string]models.StockValue
	lock   sync.Mutex
}

func newDeprecatedStockRepo() *DeprecatedStockRepo {
	return &DeprecatedStockRepo{
		stocks: make(map[string]models.StockValue),
	}
}

func (s *DeprecatedStockRepo) Add(stockValue models.StockValue) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stocks[stockValue.Label] = stockValue
}

func (s *DeprecatedStockRepo) GetStock(label string) (models.StockValue, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, found := s.stocks[label]
	if !found {
		return models.StockValue{}, fmt.Errorf("stock %s not found", label)
	}

	return v, nil
}

func (s *DeprecatedStockRepo) GetLatestStocks() (count int64, values []*models.StockValue) {
	count = int64(len(s.stocks))

	for _, v := range s.stocks {
		tmp := v
		values = append(values, &tmp)
	}

	return
}
