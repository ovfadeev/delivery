package server

import (
	"delivery/internal/app/model/repository"
	"delivery/internal/app/store"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Server {
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
	s.router.HandleFunc("/users", s.handleUsers()) // funct test query users
}

func (s *Server) handlePoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rU := fmt.Sprintf("%s", r.Header.Get("User"))
		rK := fmt.Sprintf("%s", r.Header.Get("Key"))

		if rU != "" && rK != "" {
			s.logger.Info(fmt.Sprintf("User: %s, requested points. Client IP: %s", rU, r.RemoteAddr))

			u := repository.UserRepository{s.store}
			uM, err := u.GetByLoginKey(rU, rK)
			if err != nil || uM.Id < 0 {
				s.logger.Error(err.Error())
				http.Error(w, ErrorNoLogin, http.StatusBadRequest)
			} else {
				w.Write([]byte("get points"))
			}
		} else {
			http.Error(w, ErrorNoLogin, http.StatusBadRequest)
		}
	}
}

func (s *Server) handleUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		rU := repository.UserRepository{s.store}
		u, err := rU.GetByLogin(string("ima"))
		if err != nil {
			io.WriteString(w, err.Error())
		}

		b, err := json.Marshal(u)
		if err != nil {
			return
		}
		_, err = io.WriteString(w, string(b))
		if err != nil {
			return
		}
	}
}

func (s *Server) configStore() error {
	st := store.New(s.config.Store)

	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	s.logger.Info("Database connected successful")

	return nil
}
