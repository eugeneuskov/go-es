package services

import "github.com/elastic/go-elasticsearch/v7"

type Services struct {
	Info
}

func NewServices(esClient *elasticsearch.Client) *Services {
	return &Services{
		Info: NewInfoService(esClient),
	}
}
