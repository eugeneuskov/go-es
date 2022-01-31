package services

import "github.com/elastic/go-elasticsearch/v7/esapi"

type Info interface {
	Info() (*esapi.Response, error)
}
