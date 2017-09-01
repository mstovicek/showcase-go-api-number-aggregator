package main

import (
	"github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type recoveryHandler struct {
	logger  *logrus.Logger
	handler http.Handler
}

func newSleepMiddleware(l *logrus.Logger, h http.Handler) http.Handler {
	return &recoveryHandler{
		logger:  l,
		handler: h,
	}
}

func (h *recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sleepStr := r.URL.Query().Get("sleep_ms")
	if sleepStr == "" {
		sleepStr = "0"
	}

	ms, err := strconv.Atoi(sleepStr)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"error": err,
			"path":  r.URL.Path,
		}).Error("Invalid sleep_ms parameter")
	}

	h.logger.WithFields(logrus.Fields{
		"ms":   ms,
		"path": r.URL.Path,
	}).Info("Sleep before response")

	time.Sleep(time.Duration(ms) * time.Millisecond)

	h.handler.ServeHTTP(w, r)
}
