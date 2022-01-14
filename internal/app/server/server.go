package server

import (
	"delivery/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	pkg    struct {
		logger *logrus.Logger
		router *mux.Router
		store  *store.Store
	}
}

func NewConfig(config *Config) *Server {
	return &Server{
		config: config,
		pkg: struct {
			logger *logrus.Logger
			router *mux.Router
			store  *store.Store
		}{logger: logrus.New(), router: mux.NewRouter()},
	}
}

func (s *Server) Start() error {
	if err := s.configStore(); err != nil {
		return err
	}

	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	s.pkg.logger.Info("Server started successful")

	return http.ListenAndServe(s.config.ServerAddr, s.pkg.router)
}

func (s *Server) configStore() error {
	st := store.Store{}

	if err := st.Open(s.config.DBUrl); err != nil {
		return err
	}

	s.pkg.store = &st

	s.pkg.logger.Info("Database connected successful")

	return nil
}

func (s *Server) configLogger() error {
	l, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	ft := new(logrus.TextFormatter)
	ft.FullTimestamp = true
	s.pkg.logger.Formatter = ft

	s.pkg.logger.SetLevel(l)

	return nil
}

func (s *Server) configRouter() {
	s.pkg.router.HandleFunc("/pickup", s.handlePickup())   // default method get
	s.pkg.router.HandleFunc("/courier", s.handleCourier()) // default method get
}
