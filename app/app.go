package app

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/eugeneuskov/log-service/api"
	"github.com/eugeneuskov/log-service/config"
	"github.com/eugeneuskov/log-service/server"
	"github.com/eugeneuskov/log-service/services"
	"log"
)

type Application struct {
	config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{config: config}
}

func (app *Application) RunApp() {
	es, err := app.esConnect()
	if err != nil {
		log.Fatalf("Failed to connect to elasticsearch: %s\n", err.Error())
		return
	}

	go func() {
		if err := server.NewHttpServer(
			app.config.App.ServicePort,
			api.NewHandler(
				services.NewServices(es),
				app.config.App.HandlerMode,
			).InitRoutes(),
		).Run(); err != nil {
			log.Fatalf("Error occured while runnig HTTP server %s\n", err.Error())
		}
	}()
}

func (app *Application) esConnect() (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{app.config.Address},
	})
}
