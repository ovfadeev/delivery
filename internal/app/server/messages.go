package server

import "fmt"

var (
	errorNoLogin = "Access danied"
	errorMethod  = "Error http method"
	reqSuccess   = "User: %s. Request successful. Client IP: %s. URL: %s. Method: %s"
	reqFail      = "User: %s, Request error auth. Client IP: %s. URL: %s. Method: %s"
)

func (s *Server) MsgErrorNoLogin() string {
	return errorNoLogin
}

func (s *Server) MsgErrorMethod() string {
	return errorMethod
}

func (s *Server) MsgReqSuccess(user string, ip string, url string, method string) string {
	return fmt.Sprintf(reqSuccess, user, ip, url, method)
}

func (s *Server) MsgReqFail(user string, ip string, url string, method string) string {
	return fmt.Sprintf(reqFail, user, ip, url, method)
}
