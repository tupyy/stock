package crawler

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

type Scheduler struct {
	crawler *Crawler
}

func NewScheduler(crawler *Crawler) *Scheduler {
	return &Scheduler{crawler}
}

func (s *Scheduler) Run(ctx context.Context) {
	log.Info("start scheduler")
	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				now := time.Now()
				if isStockMarketOpened(now) && !s.crawler.IsRunning() {
					log.Info("stock market opend. start crawling")
					s.crawler.Start(ctx)
				} else if !isStockMarketOpened(now) && s.crawler.IsRunning() {
					log.Info("stock market closed. stop crawling")
					s.crawler.Stop()
				}
			}
		}
	}()
}

func isStockMarketOpened(n time.Time) bool {
	if n.Weekday() == time.Saturday || n.Weekday() == time.Sunday {
		return false
	}

	if n.Hour() < 9 || n.Hour() > 17 {
		return false
	}

	return true
}
