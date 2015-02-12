package main

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/codegangsta/martini"
)

var port = 0
var responseMsg = ""
var ctr = 0
//------------------------------------------------------------------------
//----------------------- PORT -------------------------------------------
//------------------------------------------------------------------------
func postPort (enc Encoder, parms martini.Params) (int, string) {
        newport, err := strconv.Atoi(parms["portno"])
	
	if port != 0 {
		return http.StatusConflict, fmt.Sprintf("Already using portnumber %d\n", port)
	} else if err != nil || newport > 65535{
		return http.StatusRequestEntityTooLarge, Must(enc.Encode(
                        NewError(ErrCodeNotExist, fmt.Sprintf("Illegal portnumber %s\n", parms["portno"]))))
	} else {
		port = newport
		return http.StatusCreated, fmt.Sprintf("Using portnumber  %d \n", port)
	}
}

func getPort (enc Encoder, parms martini.Params) (int, string) {
	if port == 0 {
		return http.StatusPreconditionFailed, fmt.Sprintf("Must be created first\n")
	}
	return http.StatusOK, fmt.Sprintf("Using portnumber  %d \n", port)
}



func putPort (enc Encoder, parms martini.Params) (int, string) {
        newport, err := strconv.Atoi(parms["portno"])
	
	if port == 0 {
		return http.StatusPreconditionFailed, fmt.Sprintf("Must be created first\n")
	} else if err != nil || newport > 65535{
		return http.StatusRequestEntityTooLarge, Must(enc.Encode(
                        NewError(ErrCodeNotExist, fmt.Sprintf("Illegal portnumber %s\n", parms["portno"]))))
	} else {
		port = newport
		return http.StatusOK, fmt.Sprintf("Using portnumber  %d \n", port)
	}
}

func delPort (enc Encoder, parms martini.Params) (int, string) {
	if port==0 {
		return http.StatusNotFound, fmt.Sprintf("Port does not exist\n")
	}
	port = 0
	return http.StatusNoContent, fmt.Sprintf("Deleted port\n")
}


//------------------------------------------------------------------------
//----------------------- Response String---------------------------------
//------------------------------------------------------------------------

func postResponse (enc Encoder, parms martini.Params) (int, string) {
	
	if responseMsg == "" {
		responseMsg = parms["string"]	 
		fmt.Printf("Message is %s\n\r", responseMsg)
		return http.StatusCreated, responseMsg
	}
	
	return http.StatusConflict, fmt.Sprintf("The string is already set to '%s' use HTTP PUT to update", responseMsg)
}

func getResponse () (int, string) {
	ctr++
	return http.StatusOK, fmt.Sprintf("%s", responseMsg)
//	return http.StatusOK, fmt.Sprintf("%s  %d", responseMsg,ctr)
}

func putResponse (enc Encoder, parms martini.Params) (int, string) {
	
	// if responseMsg != "" {
	// 	responseMsg = parms["string"]	 
	// 	fmt.Printf("Message is %s\n\r", responseMsg)
	// 	return http.StatusCreated, responseMsg
	// }
	if responseMsg != "" {
		// <a href="http://www.yahoo.com">here</a>
		responseMsg = parms["string"]	 
// <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"
//
//		a1 := "<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n" 
//		a2 := "<meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\"/>\n</head>\n<body>\n<h1>"
//		a3 := "</h1>\n</body>\n</html>\n"
//		dispMsg := a1 + a2 + responseMsg + a3
//		dispMsg := "<!DOCTYPE html >\n<html>\n<body>\n<a href=\"http://www." + responseMsg + "\"></a>\n</body>\n</html>\n"
		fmt.Printf("pMsg= %s\n\r", responseMsg)
//		responseMsg =dispMsg
		return http.StatusCreated, responseMsg
	}
	
	return http.StatusConflict, fmt.Sprintf("The string is not created to '%s' use HTTP POST first", responseMsg)
}

func delResponse (enc Encoder, parms martini.Params) (int, string) {
	
	if responseMsg == "" {
		return http.StatusNotFound, fmt.Sprintf("Already empty\n")
		return http.StatusCreated, responseMsg
	}
	responseMsg = ""
	return http.StatusNoContent, fmt.Sprintf("Messages deleted")
}

