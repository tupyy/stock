package crawler

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/tupyy/stock/internal/config"
	"github.com/tupyy/stock/models"
)

// Repo is the interface of a stock repo
type Repo interface {
	Add(value models.StockValue)
}

type Crawler struct {
	client *http.Client

	// list of companies labels
	companies []string

	defaultCanCrawl bool

	workers map[string]*crawlWorker

	output chan models.StockValue

	stockRepo Repo

	crawlersCancelFunc context.CancelFunc

	crawlerContext context.Context
}

func NewCawler(stockRepo Repo) *Crawler {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Crawler{
		client:    &http.Client{Transport: tr},
		output:    make(chan models.StockValue),
		workers:   make(map[string]*crawlWorker),
		stockRepo: stockRepo,
	}
}
func (c *Crawler) Start(ctx context.Context) {
	c.crawlerContext, c.crawlersCancelFunc = context.WithCancel(ctx)

	go func() {
		for {
			select {
			case <-c.crawlerContext.Done():
				log.Info("context closed")
				return
			case v := <-c.output:
				log.Debug("stock saved")
				c.stockRepo.Add(v)
			}
		}
	}()

	t := config.GetCrawlPeriod()

	companies := config.GetCompanies()
	if len(companies) == 0 {
		log.Warn("no companies defined")
	}

	for _, company := range companies {
		log.Infof("starting crawler for %s", company)
		w := newcrawlWorker(c.output, t)
		go w.Run(c.crawlerContext, c.client, company)
		c.workers[company] = w
	}
}

func (c *Crawler) Stop() {
	for _, company := range c.Companies() {
		c.DeleteCompany(company)
	}
	c.crawlersCancelFunc()
}

func (c *Crawler) IsRunning() bool {
	return len(c.workers) != 0
}

func (c *Crawler) AddCompany(company string) {
	w := newcrawlWorker(c.output, 2*time.Second)
	go w.Run(c.crawlerContext, c.client, company)
	c.workers[company] = w
}

func (c *Crawler) DeleteCompany(company string) error {
	if w, ok := c.workers[company]; ok {
		log.Infof("remove worker for company '%s'", company)
		w.Shutdown()
		w = nil
		delete(c.workers, company)
		return nil
	}

	return errors.New("company not found")
}

func (c *Crawler) Companies() []string {
	var companies []string
	for k := range c.workers {
		companies = append(companies, k)
	}

	return companies
}
