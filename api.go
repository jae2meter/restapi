package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/codegangsta/martini"
)

var portno = "66666\n"
var portnoi = 0

func SetPort (enc Encoder, parms martini.Params) (int, string) {
        port, err := strconv.Atoi(parms["id"])
	if err != nil || port > 65535{
		return http.StatusNotFound, Must(enc.Encode(
                        NewError(ErrCodeNotExist, fmt.Sprintf("Illegal portnumber %s ", parms["id"]))))
	}
	return http.StatusOK, "helt OK"
}






// func GetPort(r *http.Request, enc Encoder) string {
// 	// Get the query string arguments, if any
// 	qs := r.URL.Query()
// 	wasport := qs.Get("port")
// //	portnoi, err := strconv.Atoi(portno)
// //	if err != nil {
// //		portnoi = 0
// //	}
// 	if wasport == "" {
// 		return portno
// 	}
// 	return wasport
// }


