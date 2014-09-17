package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/codegangsta/martini"
)

var port = 0

func SetPort (enc Encoder, parms martini.Params) (int, string) {
        newport, err := strconv.Atoi(parms["portno"])
	
	if port != 0 {
		return http.StatusConflict, fmt.Sprintf("Already using portnumber %d\n", port)
	} else if err != nil || port > 65535{
		return http.StatusRequestEntityTooLarge, Must(enc.Encode(
                        NewError(ErrCodeNotExist, fmt.Sprintf("Illegal portnumber %s\n", parms["portno"]))))
	} else {
		port = newport
		return http.StatusCreated, fmt.Sprintf("Using portnumber  %d \n", port)
	}
}


func SetResponse (enc Encoder, parms martini.Params) (int, string) {
	respMsg := parms["restMsg"]
	if len (respMsg) < 81 {

		return http.StatusCreated, fmt.Sprintf("Using '%s' for responses\n", respMsg)
	}

	return 413, fmt.Sprintf("The string is too long, %d characters, max 80 allowed.\n", len(respMsg))


}




