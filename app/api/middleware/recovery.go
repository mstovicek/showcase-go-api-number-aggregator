package middleware

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
)

type recoveryHandler struct {
	logger  *log.Logger
	handler http.Handler
}

func NewRecovery(l *log.Logger, h http.Handler) http.Handler {
	return &recoveryHandler{
		logger:  l,
		handler: h,
	}
}

func (h *recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		// recover allows you to continue execution in case of panic
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal Server Error")

			h.logger.WithFields(log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"error":  err,
			}).Error("Internal Server Error")
		}
	}()

	h.handler.ServeHTTP(w, r)
}
