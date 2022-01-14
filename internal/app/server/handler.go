package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handlePickup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, login, _ := s.auth(r.Header)

		if res {
			s.pkg.logger.Info(s.msgReqPointsSuccess(login, r.RemoteAddr))

			p, v := r.URL.Query().Get("param"), r.URL.Query().Get("value")

			l, err := s.getPoints(p, v)
			if err != nil {
				s.pkg.logger.Error(fmt.Sprintf("%s - %s, %s", p, v, err.Error()))
			}

			w.Write(l)
		} else {
			s.pkg.logger.Error(s.msgReqPointsFail(login, r.RemoteAddr))
			http.Error(w, s.msgErrorNoLogin(), http.StatusUnauthorized)
		}
	}
}

func (s *Server) handleCourier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, login, _ := s.auth(r.Header)

		if res {
			s.pkg.logger.Info(s.msgReqCourierSuccess(login, r.RemoteAddr))

			p, v := r.URL.Query().Get("param"), r.URL.Query().Get("value")

			l, err := s.getCourier(p, v)
			if err != nil {
				s.pkg.logger.Error(err.Error())
			}

			w.Write(l)
		} else {
			s.pkg.logger.Error(s.msgReqCourierFail(login, r.RemoteAddr))
			http.Error(w, s.msgErrorNoLogin(), http.StatusUnauthorized)
		}
	}
}
