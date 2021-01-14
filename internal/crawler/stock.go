package crawler

import (
	"fmt"
	"sync"

	"github.com/tupyy/stock/models"
)

type StockContainer struct {
	stocks map[string]StockValue
	lock   sync.Mutex
}

func newStocks() *StockContainer {
	return &StockContainer{
		stocks: make(map[string]StockValue),
	}
}

func (s *StockContainer) addStockValue(stockValue StockValue) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.stocks[stockValue.label] = stockValue
}

func (s *StockContainer) getStock(label string) (StockValue, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	v, found := s.stocks[label]
	if !found {
		return StockValue{}, fmt.Errorf("stock %s not found", label)
	}

	return v, nil
}
