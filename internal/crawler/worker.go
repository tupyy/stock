package crawler

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tupyy/stock/models"
)

type CrawlWorker struct {
	done chan struct{}

	// return true if it ok to crawl
	c canCrawl

	// interval of crawling
	tick time.Duration

	// output channel
	output chan models.StockValue
}

func NewCrawlWorker(output chan models.StockValue, tick time.Duration, c canCrawl) *CrawlWorker {
	return &CrawlWorker{
		done:   make(chan struct{}),
		c:      c,
		tick:   tick,
		output: output,
	}
}

func (c *CrawlWorker) Shutdown() {
	c.done <- struct{}{}
}

func (c *CrawlWorker) Run(ctx context.Context, client *http.Client, company string) {
	log.Infof("start crawling for '%s'", company)
	for {
		select {
		case <-ctx.Done():
			log.Info("exit from crawler")
			return
		case <-c.done:
			log.Info("exit from crawler")
			return
		case <-time.After(c.tick):
			if c.c() {
				val, err := getStock(client, company)
				if err != nil {
					log.Errorf("error getting stock %v", err)
				} else {
					log.Debugf("stock :%+v", val)
					c.output <- val
				}
			} else {
				log.Debug("cannot crawl")
			}
		}

	}
}
