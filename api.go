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
		return http.StatusConflict, fmt.Sprintf("Already using portnumber %d\n", newport)
	} else if err != nil || newport > 65535{
		return http.StatusRequestEntityTooLarge, Must(enc.Encode(
                        NewError(ErrCodeNotExist, fmt.Sprintf("Illegal portnumber %s\n", parms["portno"]))))
	} else {
		port = newport
		return http.StatusCreated, fmt.Sprintf("Using portnumber  %d \n", port)
	}
}

var responseMsg = ""

// func postResponser (r *http.Request, enc Encoder, parms martini.Params) (int, string) {
// 	id= parms["id"]

// }

 func postResponse (enc Encoder, parms martini.Params) (int, string) {

	 if responseMsg == "" {
		 responseMsg = parms["response"]	 
		 return http.StatusCreated, responseMsg
	 }
	 
	 return http.StatusConflict, fmt.Sprintf("The string is already set to '%s' use HTTP PUT to update", responseMsg)
 }

func getResponse () (int, string) {
	return http.StatusOK, fmt.Sprintf("%s", responseMsg)
}
	 
