package crawler

import (
	"context"
	"net/http"
    "fmt"
	"time"

	log "github.com/golang/glog"
)

const (
	api = "https://www.boursorama.com/bourse/action/graph/ws/UpdateCharts"
)

var (
	client *http.Client

	// list of companies labels
	companies []string
)

func Start(ctx context.Context, companies []string) *StockContainer {
	client = &http.Client{}
	stocks := newStocks()

	output := make(chan StockValue)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-output:
				stocks.addStockValue(v)
			}
		}
	}()

	t := 2 * time.Second
	for _, s := range companies {
		u := createUrl(s)

		log.V(1).Infof("starting crawler for %s", s)
		go crawl(ctx, client, u, output, t)
	}

	return stocks

}

func crawl(ctx context.Context, client *http.Client, url string, output chan StockValue, tick time.Duration) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(tick):
			val, err := getStock(client, url)
			if err != nil {
				log.V(1).Error(fmt.Sprintf("error getting stock %v", err))
			} else {
				log.V(2).Info(fmt.Sprintf("stock :%+v", val))
				output <- val
			}

		}
	}

}
