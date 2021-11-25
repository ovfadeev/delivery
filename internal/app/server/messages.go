package server

import "fmt"

var (
	ErrorNoLogin     = "Access danied"
	ReqPointsSuccess = "User: %s, requested points successful. Client IP: %s"
	ReqPointsFail = "User: %s, requested points error auth. Client IP: %s"
)

func msgErrorNoLogin() string {
	return ErrorNoLogin
}

func msgReqPointsSuccess(user string, ip string) string {
	return fmt.Sprintf(ReqPointsSuccess, user, ip)
}

func msgReqPointsFail(user string, ip string) string {
	return fmt.Sprintf(ReqPointsFail, user, ip)
}
