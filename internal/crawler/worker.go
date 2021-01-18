package crawler

import (
	"context"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tupyy/stock/models"
)

type crawlWorker struct {
	done chan struct{}

	// interval of crawling
	tick time.Duration

	// output channel
	output chan models.StockValue
}

func newcrawlWorker(output chan models.StockValue, tick time.Duration) *crawlWorker {
	return &crawlWorker{
		done:   make(chan struct{}),
		tick:   tick,
		output: output,
	}
}

func (c *crawlWorker) Shutdown() {
	c.done <- struct{}{}
}

func (c *crawlWorker) Run(ctx context.Context, client *http.Client, company string) {
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
			val, err := getStock(client, company)
			if err != nil {
				log.Errorf("error getting stock %v", err)
			} else {
				log.Debugf("stock :%+v", val)
				c.output <- val
			}
		}

	}
}
