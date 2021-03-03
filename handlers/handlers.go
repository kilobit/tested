/* Copyright 2021 Kilobit Labs Inc. a*/

// Quickly build http handlers that return a set http.StatusCode.
//
package handlers

import _ "fmt"
import _ "errors"
import "net/http"

func init() {
	OkHandler = StatusCodeHandler(http.StatusOK)
}

var OkHandler http.Handler

var NotFoundHandler http.Handler = http.NotFoundHandler()

func StatusCodeHandler(statusCode int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(http.StatusText(statusCode)))
	})
}
