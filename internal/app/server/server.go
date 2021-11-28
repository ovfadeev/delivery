package server

import (
	"delivery/internal/app/store"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
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
	st := store.New(s.config.Store)

	if err := st.Open(s.config.DBUrl); err != nil {
		return err
	}

	s.pkg.store = st

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
	s.pkg.router.HandleFunc("/points", s.handlePoints())
}

func (s *Server) handlePoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hL, hK := s.getHeaderForAuth(r.Header)

		if hL != "" && hK != "" {
			e := false
			idU, err := s.pkg.store.GetUserLogin(hL, hK)
			if err != nil || idU < 0 {
				s.pkg.logger.Error(err.Error())
				http.Error(w, s.msgErrorNoLogin(), http.StatusUnauthorized)
			} else {
				if z := r.URL.Query().Get("zip"); z != "" {
					l, err := s.pkg.store.GetPointsFromZip(z)
					if err != nil {
						e = true
						s.pkg.logger.Error(fmt.Sprintf("Request zip: %s. ", z), err.Error())
					}
					w.Write(l)
				} else if c := r.URL.Query().Get("city"); c != "" {
					l, err := s.pkg.store.GetPointsFromCity(c)
					if err != nil {
						e = true
						s.pkg.logger.Error(fmt.Sprintf("Request sity: %s. ", c), err.Error())
					}
					w.Write(l)
				}
			}
			if !e {
				s.pkg.logger.Info(s.msgReqPointsSuccess(hL, r.RemoteAddr))
			}
		} else {
			s.pkg.logger.Error(s.msgReqPointsFail(hL, r.RemoteAddr))
			http.Error(w, s.msgErrorNoLogin(), http.StatusUnauthorized)
		}
	}
}
