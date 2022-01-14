package server

import (
	"fmt"
	"net/http"
)

func (s *Server) auth(header http.Header) (bool, string, error) {
	login := fmt.Sprintf("%s", header.Get("Login"))
	key := fmt.Sprintf("%s", header.Get("Key"))

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

func (s *Server) getPoints(param string, value string) ([]byte, error) {
	if param == "zip" && len(value) > 0 {
		l, err := s.pkg.store.GetPointsFromZip(value)
		return l, err
	}
	if param == "city" && len(value) > 0 {
		l, err := s.pkg.store.GetPointsFromCity(value)
		return l, err
	}

	return []byte(""), nil
}

func (s *Server) getCourier(param string, value string) ([]byte, error) {

	return []byte(""), nil
}
