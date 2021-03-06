// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-redis/redis"

	"github.com/tupyy/stock/internal/config"
	"github.com/tupyy/stock/internal/crawler"
	"github.com/tupyy/stock/internal/repo"
	"github.com/tupyy/stock/models"
	"github.com/tupyy/stock/restapi/operations"

	log "github.com/sirupsen/logrus"
)

//go:generate swagger generate server --target ../../stock-crawler --name StockService --spec ../target/swagger.yaml --principal interface{}

type configurationFlags struct {
	ConfFile string `short:"c" long:"conf" description:"Path to configuration file" value-name:"FILE"`
}

var confFlags configurationFlags
var c *crawler.Crawler

func configureFlags(api *operations.StockServiceAPI) {

	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Configuration Options",
			Options:          &confFlags,
		},
	}
}

func configureAPI(api *operations.StockServiceAPI) http.Handler {
	configureGlobal()

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	ctx, cancel := context.WithCancel(context.Background())

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisUrl(), // use default Addr
		Password: "",                   // no password set
		DB:       0,                    // use default DB
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("connected to redis %s", config.GetRedisUrl())

	redisRepo := repo.NewStockRedisRepo(rdb)
	c := crawler.NewCawler(redisRepo)
	if config.UseScheduler() {
		scheduler := crawler.NewScheduler(c)
		scheduler.Run(ctx)
	} else {
		c.Start(ctx)
	}

	if api.GetStockHandler == nil {
		api.GetStockHandler = operations.GetStockHandlerFunc(func(params operations.GetStockParams) middleware.Responder {
			payload := models.Stock{}
			if params.Label != nil {
				s, err := redisRepo.GetStock(*params.Label)
				if err != nil {
					return operations.NewGetStockNotFound()
				}

				payload.Count = 1
				payload.Values = append(payload.Values, &s)
				return operations.NewGetStockOK().WithPayload(&payload)
			}

			payload.Count, payload.Values = redisRepo.GetLatestStocks()
			if payload.Count == 0 {
				return &operations.GetStockOK{}
			}
			return operations.NewGetStockOK().WithPayload(&payload)
		})
	}
	if api.GetCompaniesHandler == nil {
		api.GetCompaniesHandler = operations.GetCompaniesHandlerFunc(func(params operations.GetCompaniesParams) middleware.Responder {
			payload := &operations.GetCompaniesOKBody{Companies: c.Companies()}

			return operations.NewGetCompaniesOK().WithPayload(payload)
		})
	}
	if api.PostStockCompanyHandler == nil {
		api.PostStockCompanyHandler = operations.PostStockCompanyHandlerFunc(func(params operations.PostStockCompanyParams) middleware.Responder {
			c.AddCompany(params.Company)
			return operations.NewPostStockCompanyCreated()
		})
	}

	if api.DeleteStockCompanyHandler == nil {
		api.DeleteStockCompanyHandler = operations.DeleteStockCompanyHandlerFunc(func(params operations.DeleteStockCompanyParams) middleware.Responder {
			err := c.DeleteCompany(params.Company)
			if err != nil {
				return operations.NewDeleteStockCompanyNotFound()
			}

			return operations.NewDeleteStockCompanyOK()
		})
	}

	if api.GetStocksCompanyHandler == nil {
		api.GetStocksCompanyHandler = operations.GetStocksCompanyHandlerFunc(func(params operations.GetStocksCompanyParams) middleware.Responder {
			payload, err := redisRepo.GetStocks(params.Company)
			if err != nil {
				return operations.NewGetStocksCompanyNotFound()
			}

			return operations.NewGetStocksCompanyOK().WithPayload(&payload)

		})
	}

	api.PreServerShutdown = func() {
		cancel()
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

func configureGlobal() {
	// Set log output
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{})

	config.ParseConfiguration(confFlags.ConfFile)

	log.SetLevel(log.InfoLevel)
	if os.Getenv("DEBUG") == "1" {
		log.SetLevel(log.DebugLevel)
	}
}
