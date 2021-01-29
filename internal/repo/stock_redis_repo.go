package repo

import (
	"errors"
	"strconv"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/tupyy/stock/models"
)

var (
	ErrStockNotFound = errors.New("stock not found")
)

type StockRedisRepo struct {
	stockRepo *DeprecatedStockRepo

	redisClient *redis.Client
}

func NewStockRedisRepo(rdc *redis.Client) *StockRedisRepo {
	stockRepo := newDeprecatedStockRepo()

	return &StockRedisRepo{
		stockRepo:   stockRepo,
		redisClient: rdc,
	}
}

func (s *StockRedisRepo) Add(stockValue models.StockValue) {
	if lastVal, err := s.stockRepo.GetStock(stockValue.Label); err == nil {
		if lastVal.Value != stockValue.Value {
			log.WithField("stock", stockValue).Debug("push value to redis")
			s.redisClient.LPush(stockValue.Label, stockValue.Value)
		}
	}

	s.stockRepo.Add(stockValue)
}

func (s *StockRedisRepo) GetStock(label string) (models.StockValue, error) {
	return s.stockRepo.GetStock(label)
}

func (s *StockRedisRepo) GetStocks(label string) (models.StockValues, error) {
	lenCmd := s.redisClient.LLen(label)
	if _, err := lenCmd.Result(); err != nil {
		return models.StockValues{}, ErrStockNotFound
	}

	count := lenCmd.Val()
	values := make([]float64, 0, count)
	if count == 0 {
		return models.StockValues{}, ErrStockNotFound
	}

	strSliceCmd := s.redisClient.LRange(label, 0, count)
	for _, s := range strSliceCmd.Val() {
		if val, err := strconv.ParseFloat(s, 64); err == nil {
			values = append(values, val)
		}
	}

	return redisToStockValues(count, label, values), nil
}

func (s *StockRedisRepo) GetLatestStocks() (count int64, values []*models.StockValue) {
	return s.stockRepo.GetLatestStocks()
}
