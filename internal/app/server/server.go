package server

import (
	"delivery/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewConfig(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	if err := s.configStore(); err != nil {
		return err
	}

	s.logger.Info("Server started successful")

	return http.ListenAndServe(s.config.ServerAddr, s.router)
}

func (s *Server) configLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	ft := new(logrus.TextFormatter)
	ft.FullTimestamp = true
	s.logger.Formatter = ft

	s.logger.SetLevel(level)

	return nil
}

func (s *Server) configRouter() {
	s.router.HandleFunc("/points", s.handlePoints())
}

func (s *Server) handlePoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hL, hK := getHeaderForAuth(r.Header)

		if hL != "" && hK != "" {
			s.logger.Info(msgReqPointsSuccess(hL, r.RemoteAddr))

			uM, err := s.store.User().GetByLoginKey(hL, hK)
			if err != nil || uM.Id < 0 {
				s.logger.Error(err.Error())
				http.Error(w, msgErrorNoLogin(), http.StatusBadRequest)
			} else {
				w.Write([]byte("get points"))
			}
		} else {
			s.logger.Error(msgReqPointsFail(hL, r.RemoteAddr))
			http.Error(w, msgErrorNoLogin(), http.StatusBadRequest)
		}
	}
}

func (s *Server) configStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(s.config.DBUrl); err != nil {
		return err
	}

	s.store = st

	s.logger.Info("Database connected successful")

	return nil
}
