package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handlePickup() http.HandlerFunc {
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

func (s *Server) handleCourier() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
