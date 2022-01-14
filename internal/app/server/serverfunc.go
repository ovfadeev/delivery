package server

import (
	"fmt"
	"net/http"
	"net/url"
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

func (s *Server) getPoints(query url.Values) ([]byte, error) {
	if query["param"][0] == "zip" && len(query["value"][0]) > 0 {
		l, err := s.pkg.store.GetPointsFromZip(query["value"][0])
		return l, err
	}
	if query["param"][0] == "city" && len(query["value"][0]) > 0 {
		l, err := s.pkg.store.GetPointsFromCity(query["value"][0])
		return l, err
	}

	return []byte(""), nil
}

func (s *Server) getCourier(param string, value string) ([]byte, error) {

	return []byte(""), nil
}
