package main

import (
//	"fmt"
	"net/http"
//	"strconv"

//	"github.com/codegangsta/martini"
)

var portno = "66666\n"
var portnoi = 0


func GetPort(r *http.Request, enc Encoder) string {
	// Get the query string arguments, if any
	qs := r.URL.Query()
	wasport := qs.Get("port")
//	portnoi, err := strconv.Atoi(portno)
//	if err != nil {
//		portnoi = 0
//	}
	if wasport == "" {
		return portno
	}
	return "deFel"
}


