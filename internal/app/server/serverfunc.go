package server

import (
	"fmt"
	"net/http"
)

func getHeaderForAuth(header http.Header) (string, string) {
	return fmt.Sprintf("%s", header.Get("Login")), fmt.Sprintf("%s", header.Get("Key"))
}
