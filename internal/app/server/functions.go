package server

import (
	"net/http"
	"net/url"
)

func (s *Server) auth(header http.Header) (bool, string, error) {
	login := header.Get("Login") // return string
	key := header.Get("Key")     // return string

	if login != "" && key != "" {
		idU, err := s.pkg.store.GetUserLogin(login, key)
		if err != nil || idU < 0 {
			return false, login, err
		} else {
			return true, login, nil
		}
	}
	return false, "\"\"", nil
}

func (s *Server) getData() {

}

func (s *Server) getPoints(query url.Values) ([]byte, error) {
	if query["zip"][0] != "" {
		s.pkg.logger.Info(query)
		l, err := s.pkg.store.GetPointsFromZip(query["zip"][0])
		return l, err
	} else if query["city"][0] != "" {
		l, err := s.pkg.store.GetPointsFromZip(query["city"][0])
		return l, err
	}

	return []byte(""), nil
}

func (s *Server) getCourier(param string, value string) ([]byte, error) {

	return []byte(""), nil
}
