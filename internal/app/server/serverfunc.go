package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Server) getHeaderForAuth(header http.Header) (string, string) {
	return fmt.Sprintf("%s", header.Get("Login")), fmt.Sprintf("%s", header.Get("Key"))
}

func (s *Server) getPointsFromCity(city string) ([]byte, error) {
	p := s.store.Point().GetByCity(city)
	return json.Marshal(p)
}
