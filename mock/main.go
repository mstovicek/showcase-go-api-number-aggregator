package main

import "github.com/Sirupsen/logrus"

func main() {
	apiServer := newServer(
		logrus.New(),
	)
	apiServer.Run()
}
