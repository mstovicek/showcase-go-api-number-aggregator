package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	logger *logrus.Logger
}

func newServer(l *logrus.Logger) *server {
	return &server{
		logger: l,
	}
}

func (s *server) Run() {
	router := mux.NewRouter()

	router.Handle(
		"/ten/",
		newTen(),
	).Methods(http.MethodGet)

	router.Handle(
		"/even/",
		newEven(),
	).Methods(http.MethodGet)

	router.Handle(
		"/odd/",
		newOdd(),
	).Methods(http.MethodGet)

	router.Handle(
		"/fibo/",
		newFibo(),
	).Methods(http.MethodGet)

	router.Handle(
		"/primes/",
		newPrimes(),
	).Methods(http.MethodGet)

	router.Handle(
		"/wrong/",
		newWrong(),
	).Methods(http.MethodGet)

	http.Handle("/", router)

	h := newSleepMiddleware(
		s.logger,
		router,
	)

	s.logger.Fatal(http.ListenAndServe(":8090", h))
}
