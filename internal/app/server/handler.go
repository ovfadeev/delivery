package server

import (
	"net/http"
)

func (s *Server) HandlePickup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, login, _ := s.Auth(r.Header)

		if r.Method != "GET" {
			s.pkg.logger.Info(s.MsgReqFail(login, r.RemoteAddr, r.RequestURI, r.Method))
			s.pkg.logger.Error(s.MsgErrorMethod())
			http.Error(w, s.MsgErrorMethod(), http.StatusUnauthorized)
			return
		}

		if res {
			s.pkg.logger.Info(s.MsgReqSuccess(login, r.RemoteAddr, r.RequestURI, r.Method))

			l, err := s.GetPoints(r.URL.Query())
			if err != nil {
				s.pkg.logger.Error(err.Error())
			}

			w.Write(l)
		} else {
			s.pkg.logger.Info(s.MsgReqFail(login, r.RemoteAddr, r.RequestURI, r.Method))
			http.Error(w, s.MsgErrorNoLogin(), http.StatusUnauthorized)
		}
	}
}

func (s *Server) HandleCourier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
