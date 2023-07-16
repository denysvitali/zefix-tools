package processor

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/denysvitali/documents-indexer/pkg"
	"github.com/denysvitali/go-datesfinder"
	"github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Processor struct {
	index    string
	logger   *logrus.Logger
	osClient *opensearch.Client
}

func New(osAddr string, osUsername string, osPassword string, skipTls bool, osIndex string) (*Processor, error) {
	osClient, err := opensearch.NewClient(
		opensearch.Config{
			Addresses: []string{osAddr},
			Username:  osUsername,
			Password:  osPassword,
			Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTls}},
		},
	)

	if err != nil {
		return nil, err
	}

	p := Processor{
		osClient: osClient,
		index:    osIndex,
		logger:   logrus.New(),
	}
	return &p, nil
}

func (p *Processor) SetLogger(logger *logrus.Logger) {
	p.logger = logger
}

func (p *Processor) Process() error {
	// Go through all the documents in the index and update them
	// by adding some new fields

	// 1. Get all the documents from the index
	size := 1000
	req := opensearchapi.SearchRequest{
		Index: []string{p.index},
		Sort:  []string{"date:asc"},
		Size:  &size,
	}

	res, err := req.Do(context.Background(), p.osClient)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// Parse response as JSON
	var result OpensearchResult[indexer.Document]
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&result)
	if err != nil {
		return err
	}

	// 2. Iterate over the documents
	for _, hit := range result.Hits.Hits {
		p.processText(hit.Source.Text)
	}

	return nil
}

func (p *Processor) processText(text string) {
	// Find dates
	dates, errors := datesfinder.FindDates(text)
	if len(errors) != 0 {
		p.logger.Warnf("found %d errors while parsing dates", len(errors))
		for _, err := range errors {
			p.logger.Warnf("error: %s", err)
		}
	}
	p.logger.Infof("found %d dates", len(dates))
}
