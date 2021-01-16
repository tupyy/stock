package crawler

import (
	"context"
	"crypto/tls"
	"errors"
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

	workers map[string]*CrawlWorker

	output chan models.StockValue

	parentContext context.Context
)

func Start(ctx context.Context, companies []string) *StockContainer {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	stocks := newStocks()

	parentContext = ctx
	output = make(chan models.StockValue)
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Info("context closed")
				return
			case v := <-output:
				log.Debug("stock saved")
				stocks.addStockValue(v)
			}
		}
	}()

	t := 2 * time.Second
	workers = make(map[string]*CrawlWorker)
	for _, c := range companies {
		log.Infof("starting crawler for %s", c)
		w := NewCrawlWorker(output, t, createCanCrawl())
		go w.Run(parentContext, client, c)
		workers[c] = w
	}

	return stocks
}

func AddCompany(company string) {
	w := NewCrawlWorker(output, 2*time.Second, createCanCrawl())
	go w.Run(parentContext, client, company)
	workers[company] = w
}

func DeleteCompany(company string) error {
	if w, ok := workers[company]; ok {
		w.Shutdown()
		w = nil
		delete(workers, company)
		return nil
	}

	return errors.New("company not found")
}

// return function to schedule crawling.
func createCanCrawl() canCrawl {
	return func() bool {
		now := time.Now()

		// dont crawl in weekends
		if now.Day() > 4 {
			return false
		}

		// crawl between 9 - 18
		if now.Hour() >= 9 && now.Hour() < 18 {
			return true
		}
		return false
	}
}
