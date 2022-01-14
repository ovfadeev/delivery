package server

import "fmt"

var (
	errorNoLogin = "Access danied"
	errorMethod  = "Error http method"
	reqSuccess   = "User: %s. Request successful. Client IP: %s. URL: %s. Method: %s"
	reqFail      = "User: %s, Request error auth. Client IP: %s. URL: %s. Method: %s"
)

func (s *Server) msgErrorNoLogin() string {
	return errorNoLogin
}

func (s *Server) msgErrorMethod() string {
	return errorMethod
}

func (s *Server) msgReqSuccess(user string, ip string, url string, method string) string {
	return fmt.Sprintf(reqSuccess, user, ip, url, method)
}

func (s *Server) msgReqFail(user string, ip string, url string, method string) string {
	return fmt.Sprintf(reqFail, user, ip, url, method)
}
