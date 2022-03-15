package server

import (
	"delivery/internal/app/delivery"
	"delivery/internal/app/store"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *Config
	pkg    struct {
		logger   *logrus.Logger
		router   *mux.Router
		store    *store.Store
		delivery *delivery.Delivery
	}
}

func NewConfig(config *Config) *Server {
	return &Server{
		config: config,
		pkg: struct {
			logger   *logrus.Logger
			router   *mux.Router
			store    *store.Store
			delivery *delivery.Delivery
		}{logger: logrus.New(), router: mux.NewRouter()},
	}
}

func (s *Server) Start() error {
	if err := s.ConfigStore(); err != nil {
		return err
	}

	if err := s.ConfigLogger(); err != nil {
		return err
	}

	if err := s.ConfigDelivery(); err != nil {
		return err
	}

	s.ConfigRouter()

	s.pkg.logger.Info("Server started successful")

	return http.ListenAndServe(s.config.ServerAddr, s.pkg.router)
}

func (s *Server) ConfigStore() error {
	st := store.Store{}

	if err := st.Open(s.config.DBUrl); err != nil {
		return err
	}

	s.pkg.store = &st

	s.pkg.logger.Info("Database connected successful")

	return nil
}

func (s *Server) ConfigLogger() error {
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

func (s *Server) ConfigDelivery() error {
	dl := delivery.Delivery{}

	dl.NewConfig(s.config.Delivery)

	s.pkg.delivery = &dl

	return nil
}

func (s *Server) ConfigRouter() {
	s.pkg.router.HandleFunc("/pickup", s.HandlePickup())   // default method get
	s.pkg.router.HandleFunc("/courier", s.HandleCourier()) // default method get
}
