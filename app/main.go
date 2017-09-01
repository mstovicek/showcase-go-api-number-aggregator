package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/mstovicek/showcase-go-api-number-aggregator/app/api"
)

func main() {
	apiServer := api.NewServer(
		logrus.New(),
	)
	apiServer.Run()
}
