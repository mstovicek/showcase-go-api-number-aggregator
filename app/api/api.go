package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/jessevdk/go-flags"
	"github.com/mstovicek/showcase-go-api-number-aggregator/app/api/handler"
	"github.com/mstovicek/showcase-go-api-number-aggregator/app/api/middleware"
	"github.com/mstovicek/showcase-go-api-number-aggregator/app/loader"
	"net/http"
)

type serverConfig struct {
	Listen        string `long:"listen" description:"Listen address" default:":8080" required:"true"`
	ReaderTimeout int    `long:"reader-timeout" description:"Timeout to read numbers" default:"490" required:"true"`
}

type Server interface {
	Run()
}

type server struct {
	logger *logrus.Logger
	config serverConfig
}

func NewServer(l *logrus.Logger) Server {
	var conf serverConfig
	if _, err := flags.NewParser(&conf, flags.HelpFlag|flags.PassDoubleDash).Parse(); err != nil {
		l.Fatalln(err)
	}

	return &server{
		logger: l,
		config: conf,
	}
}

func (s *server) Run() {
	router := mux.NewRouter()

	router.Handle(
		"/numbers/",
		handler.NewNumbers(
			loader.NewLoader(
				loader.NewURLValidator(),
				loader.NewHTTPReader(
					s.logger,
					s.config.ReaderTimeout,
				),
			),
		),
	).Methods(http.MethodGet)

	http.Handle("/", router)

	h := middleware.NewRecovery(
		s.logger,
		middleware.NewLogResponseTime(
			s.logger,
			router,
		),
	)

	s.logger.WithFields(logrus.Fields{
		"listen":         s.config.Listen,
		"reader-timeout": s.config.ReaderTimeout,
	}).Infof("Listening on the address %s", s.config.Listen)
	s.logger.Fatal(http.ListenAndServe(s.config.Listen, h))
}
