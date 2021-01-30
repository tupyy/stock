package repo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/tupyy/stock/internal/config"
	"github.com/tupyy/stock/models"
)

const (

	// format of the key where the max value of the day is stored
	// %s must be replaces with company label
	maxValueKeyFormat = "%s_max"

	// format of the key where the min value of the day is stored
	// %s must be replaces with company label
	minValueKeyFormat = "%s_min"

	// noExpiration for redis key/values
	noExpiration = 0 * time.Second
)

var (

	// ErrStockNotFound means that the company is not found in db
	ErrStockNotFound = errors.New("stock not found")

	// ErrStockNoValues means that the key exists but there are no values
	ErrStockNoValues = errors.New("stock no values")

	// ErrRedis unknown redis error
	ErrRedis = errors.New("redis error")
)

type StockRedisRepo struct {
	redisClient *redis.Client
}

func NewStockRedisRepo(rdc *redis.Client) *StockRedisRepo {
	return &StockRedisRepo{
		redisClient: rdc,
	}
}

func (s *StockRedisRepo) Add(stockValue models.StockValue) {
	company := strings.ToLower(stockValue.Label)

	if lastVal, err := s.GetStock(company); err == nil {
		if lastVal.Value != stockValue.Value {
			log.WithField("stock", stockValue).Debug("push value to redis")
			s.redisClient.LPush(company, stockValue.Value)

			// check if it is greater than max or less than min
			max, err := s.getValue(fmt.Sprintf(maxValueKeyFormat, company))
			if err == nil {
				if max < stockValue.Value {
					errSet := s.setValue(fmt.Sprintf(maxValueKeyFormat, company), max)
					if errSet != nil {
						logrus.WithError(err).Warnf("cannot set max value %f with key %s", max, company)
					}
				}
			}

			min, errMin := s.getValue(fmt.Sprintf(minValueKeyFormat, company))
			if errMin == nil {
				if min > stockValue.Value {
					errSet := s.setValue(fmt.Sprintf(minValueKeyFormat, company), min)
					if errSet != nil {
						logrus.WithError(err).Warnf("cannot set min value %f with key %s", min, company)
					}
				}
			}
		}
	}
}

func (s *StockRedisRepo) GetStock(label string) (models.StockValue, error) {
	lastVal := s.redisClient.LRange(label, -1, -1)
	valS, err := lastVal.Result()
	if err != nil {
		return models.StockValue{}, fmt.Errorf("[%w] %v", ErrRedis, err)
	}

	if len(valS) == 0 {
		return models.StockValue{}, ErrStockNoValues
	}

	val, err := strconv.ParseFloat(valS[0], 64)
	if err != nil {
		return models.StockValue{}, fmt.Errorf("[%w] bad convertion from string '%s' to float", ErrRedis, valS)
	}

	max, err := s.getValue(fmt.Sprintf(maxValueKeyFormat, label))
	if err != nil {
		return models.StockValue{}, fmt.Errorf("[%w] %v", ErrRedis, err)
	}

	min, err := s.getValue(fmt.Sprintf(minValueKeyFormat, label))
	if err != nil {
		return models.StockValue{}, fmt.Errorf("[%w] %v", ErrRedis, err)
	}

	variation := (1 - min/val) * 100

	return models.StockValue{
		Label:     label,
		Value:     val,
		Max:       max,
		Min:       min,
		Variation: variation,
	}, nil

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
	companies := config.GetCompanies()
	if len(companies) == 0 {
		return 0, []*models.StockValue{}
	}

	val := make([]*models.StockValue, 0, len(companies))

	for _, company := range companies {
		if v, err := s.GetStock(strings.ToLower(company)); err != nil {
			logrus.WithError(err).Warnf("cannot read latest value for company %s", company)
		} else {
			val = append(val, &v)
		}
	}

	return int64(len(val)), val
}

// getValue return the value of the key as float64.
func (s *StockRedisRepo) getValue(key string) (float64, error) {
	valCmd := s.redisClient.Get(key)
	if valS, err := valCmd.Result(); err != nil {
		return 0, err
	} else {
		val, err := strconv.ParseFloat(valS, 64)
		if err != nil {
			return 0, err
		}

		return val, nil
	}
}

func (s *StockRedisRepo) setValue(key string, value float64) error {
	setCmd := s.redisClient.Set(key, value, noExpiration)
	_, err := setCmd.Result()

	return err
}
