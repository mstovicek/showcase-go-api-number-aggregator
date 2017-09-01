package loader

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type NumberReader interface {
	ReadNumbers(url string) []int
}

type numbersJSON struct {
	Numbers []int `json:"numbers"`
}

type httpReader struct {
	logger  *logrus.Logger
	timeout int
}

func NewHTTPReader(l *logrus.Logger, t int) NumberReader {
	return &httpReader{
		logger:  l,
		timeout: t,
	}
}

func (reader *httpReader) ReadNumbers(url string) []int {
	startTime := time.Now()

	client := http.Client{
		Timeout: time.Millisecond * time.Duration(reader.timeout),
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		reader.logger.WithFields(logrus.Fields{
			"error": err,
			"url":   url,
		}).Error("Request creation failed")
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		reader.logger.WithFields(logrus.Fields{
			"error":   err,
			"url":     url,
			"timeout": client.Timeout,
		}).Warn("Request failed")
		return nil
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		reader.logger.WithFields(logrus.Fields{
			"error": err,
			"url":   url,
		}).Error("Cannot read response body")
		return nil
	}

	data := numbersJSON{}
	json.Unmarshal(body, &data)

	reader.logger.WithFields(logrus.Fields{
		"url":     url,
		"numbers": data.Numbers,
		"time":    time.Since(startTime),
	}).Info("Numbers loaded")

	return data.Numbers
}
