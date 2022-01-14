package server

import "fmt"

var (
	errorNoLogin = "Access danied"
	reqSuccess   = "User: %s. Request successful. Client IP: %s. URL: %s"
	reqFail      = "User: %s, Request error auth. Client IP: %s. URL: %s"
)

func (s *Server) msgErrorNoLogin() string {
	return errorNoLogin
}

func (s *Server) msgReqSuccess(user string, ip string, url string) string {
	return fmt.Sprintf(reqSuccess, user, ip, url)
}

func (s *Server) msgReqFail(user string, ip string, url string) string {
	return fmt.Sprintf(reqFail, user, ip, url)
}
