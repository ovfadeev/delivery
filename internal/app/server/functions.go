package server

import (
	"net/http"
	"net/url"
)

func (s *Server) Auth(header http.Header) (bool, string, error) {
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

func (s *Server) GetData() {

}

func (s *Server) GetPoints(query url.Values) ([]byte, error) {
	// pq := PrepareQuery(query)
	pq := query

	if pq["zip"][0] != "" {
		s.pkg.logger.Info(query)
		l, err := s.pkg.store.GetPointsFromZip(pq["zip"][0])
		return l, err
	} else if query["city"][0] != "" {
		l, err := s.pkg.store.GetPointsFromZip(pq["city"][0])
		return l, err
	}

	l, err := s.pkg.delivery.GetPoints(pq)

	return l, err
}

func (s *Server) GetCourier(query url.Values) ([]byte, error) {

	l, err := s.pkg.delivery.GetCourier(query)
	return l, err
}

// func PrepareQuery(query url.Values) {
// 	// res := map[string][string]{}

// 	// for i :=

// 	return query
// }
