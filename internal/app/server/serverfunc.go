package server

import (
	"fmt"
	"net/http"
)

func (s *Server) auth(header http.Header) (string, error) {
	login, key := fmt.Sprintf("%s", header.Get("Login")), fmt.Sprintf("%s", header.Get("Key"))

	if login != "" && key != "" {
		idU, err := s.pkg.store.GetUserLogin(login, key)
		if err != nil || idU < 0 {
			return "", err
		} else {
			return login, nil
		}
	}
	return "", nil
}

func (s *Server) getData() {

}
