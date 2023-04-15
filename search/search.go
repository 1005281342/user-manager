package search

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var client *elasticsearch.Client

func Connect() error {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}

	var err error
	client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	_, err = client.Info()
	if err != nil {
		return err
	}

	return nil
}

func GetClient() *elasticsearch.Client {
	return client
}
