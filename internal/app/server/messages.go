package server

import "fmt"

var (
	errorNoLogin     = "Access danied"
	reqPointsSuccess = "User: %s, requested points successful. Client IP: %s"
	reqPointsFail    = "User: %s, requested points error auth. Client IP: %s"
)

func (s *Server) msgErrorNoLogin() string {
	return errorNoLogin
}

func (s *Server) msgReqPointsSuccess(user string, ip string) string {
	return fmt.Sprintf(reqPointsSuccess, user, ip)
}

func (s *Server) msgReqPointsFail(user string, ip string) string {
	return fmt.Sprintf(reqPointsFail, user, ip)
}
