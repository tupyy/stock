package config

import (
	"time"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const prefix = "STOCK_SERVICE"

// Below all the different keys used to configure this service.
const (
	// base url for stock server (i.e. boursarama)
	stockServerURL = "stockServerURL"

	// redisURN
	redisURL = "redisURL"

	// list of companies
	companies = "companies"

	// crawl period
	crawlPeriod = "crawlPeriod"

	// if crawlPeriod is not defined, the default period is set to 2s
	defaultCrawlPeriod = 2 * time.Second

	// useScheduler means that the crawler will crawl only one the stock market is open
	useScheduler = "useScheduler"
)

// ParseConfiguration reads the configuration file given as parameter.
func ParseConfiguration(confFile string) {
	viper.SetDefault(crawlPeriod, defaultCrawlPeriod)

	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv() // read in environment variables that match

	if len(confFile) > 0 {
		viper.SetConfigFile(confFile)

		err := viper.ReadInConfig()
		if err != nil {
			logger.WithError(err).Errorf("failed to read config file %v", confFile)
		} else {
			logger.Infof("using config file: %v", viper.ConfigFileUsed())
		}
	}
}

func GetServerBaseUrl() string {
	return viper.GetString(stockServerURL)
}

func GetCompanies() []string {
	return viper.GetStringSlice(companies)
}

func GetRedisUrl() string {
	return viper.GetString(redisURL)
}

func UseScheduler() bool {
	return viper.GetBool(useScheduler)
}

func GetCrawlPeriod() time.Duration {
	val := viper.GetInt(crawlPeriod)
	return time.Duration(val) * time.Second
}
