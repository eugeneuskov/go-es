package services

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
)

type InfoService struct {
	esClient *elasticsearch.Client
}

func (i *InfoService) Info() (*esapi.Response, error) {
	res, err := i.esClient.Info()
	if err != nil {
		log.Printf("--- info err: %s\n", err.Error())
		return nil, err
	}

	if res.IsError() {
		log.Printf("--- is err: %s\n", res.String())
		return nil, fmt.Errorf("%s", res.String())
	}

	return res, nil
}

func NewInfoService(esClient *elasticsearch.Client) *InfoService {
	return &InfoService{esClient}
}
