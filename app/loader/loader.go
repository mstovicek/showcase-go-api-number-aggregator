package loader

import (
	"github.com/Sirupsen/logrus"
	"github.com/mstovicek/showcase-go-api-number-aggregator/app/Numbers"
	"sync"
)

type Loader interface {
	Load(urls []string) Numbers.NumberCollection
}

type loader struct {
	logger       *logrus.Logger
	urlValidator URLValidator
	reader       NumberReader
}

func NewLoader(urlValidator URLValidator, reader NumberReader) Loader {
	return &loader{
		logger:       logrus.New(),
		urlValidator: urlValidator,
		reader:       reader,
	}
}

func (loader *loader) Load(urls []string) Numbers.NumberCollection {
	var wg sync.WaitGroup

	collection := Numbers.NewNumberCollection()
	validUrls := loader.getValidUrls(urls)

	wg.Add(len(validUrls))

	for _, url := range validUrls {
		go func(url string) {
			defer wg.Done()

			loader.logger.WithFields(logrus.Fields{
				"url": url,
			}).Info("Loading numbers")

			collection.AddNumbers(loader.reader.ReadNumbers(url))
		}(url)
	}

	wg.Wait()

	return collection
}

func (loader *loader) getValidUrls(urls []string) []string {
	var validUrls []string

	for _, url := range urls {
		if !loader.urlValidator.IsValid(url) {
			loader.logger.WithFields(logrus.Fields{
				"url": url,
			}).Warn("Ignore invalid url")
			continue
		}

		validUrls = append(validUrls, url)
	}

	return validUrls
}
