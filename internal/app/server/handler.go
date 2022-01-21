package server

import (
	"net/http"
)

func (s *Server) handlePickup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, login, _ := s.auth(r.Header)

		if r.Method != "GET" {
			s.pkg.logger.Info(s.msgReqFail(login, r.RemoteAddr, r.RequestURI, r.Method))
			s.pkg.logger.Error(s.msgErrorMethod())
			http.Error(w, s.msgErrorMethod(), http.StatusUnauthorized)
			return
		}

		if res {
			s.pkg.logger.Info(s.msgReqSuccess(login, r.RemoteAddr, r.RequestURI, r.Method))

			l, err := s.getPoints(r.URL.Query())
			if err != nil {
				s.pkg.logger.Error(err.Error())
			}

			w.Write(l)
		} else {
			s.pkg.logger.Info(s.msgReqFail(login, r.RemoteAddr, r.RequestURI, r.Method))
			http.Error(w, s.msgErrorNoLogin(), http.StatusUnauthorized)
		}
	}
}

func (s *Server) handleCourier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
