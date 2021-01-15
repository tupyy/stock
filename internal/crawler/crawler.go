package crawler

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tupyy/stock/models"
)

const (
	api = "https://www.boursorama.com/bourse/action/graph/ws/UpdateCharts"
)

// return true if the crawler is allowed to crawl
// it is used to stop crawling when the stock market is closed
type canCrawl func() bool

var (
	client *http.Client

	// list of companies labels
	companies []string
)

func Start(ctx context.Context, companies []string) *StockContainer {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	stocks := newStocks()

	output := make(chan models.StockValue)
	go func() {
		select {
		case <-ctx.Done():
			log.Info("context closed")
			return
		case v := <-output:
			if v.Value != nil {
				stocks.addStockValue(v)
			}
		}
	}()

	t := 2 * time.Second
	for _, s := range companies {
		log.Infof("starting crawler for %s", s)
		go crawl(ctx, client, s, output, t, createCanCrawl())
	}

	return stocks

}

func crawl(ctx context.Context, client *http.Client, company string, output chan models.StockValue, tick time.Duration, c canCrawl) {
	for {
		select {
		case <-ctx.Done():
			log.Info("exit from crawler")
			return
		case <-time.After(tick):
			if c() {
				val, err := getStock(client, company)
				if err != nil {
					log.Errorf(fmt.Sprintf("error getting stock %v", err))
				} else {
					log.Debug(fmt.Sprintf("stock :%+v", val))
					output <- val
				}
			} else {
				log.Debug("cannot crawl")
			}
		}

	}
}

func createCanCrawl() canCrawl {
	return func() bool {
		now := time.Now()
		if now.Hour() > 9 && now.Hour() < 18 {
			return true
		}
		return false
	}
}
