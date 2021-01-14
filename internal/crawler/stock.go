package crawler

import (
	"fmt"
	"sync"

	"github.com/tupyy/stock/models"
)

type StockContainer struct {
	stocks map[string]models.StockValue
	lock   sync.Mutex
}

func newStocks() *StockContainer {
	return &StockContainer{
		stocks: make(map[string]models.StockValue),
	}
}

func (s *StockContainer) addStockValue(stockValue models.StockValue) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stocks[stockValue.Label] = stockValue
}

func (s *StockContainer) GetStock(label string) (models.StockValue, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, found := s.stocks[label]
	if !found {
		return models.StockValue{}, fmt.Errorf("stock %s not found", label)
	}

	return v, nil
}

func (s *StockContainer) GetStocks() (count int64, values []*models.StockValue) {
	count = int64(len(s.stocks))

	for _, v := range s.stocks {
		tmp := v
		values = append(values, &tmp)
	}

	return
}
