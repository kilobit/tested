/* Copyright 2020 Kilobit Labs Inc. */

package handlers_test

import _ "fmt"
import _ "errors"
import "net/http"
import "net/http/httptest"

import "testing"
import "kilobit.ca/go/tested/assert"

import "kilobit.ca/go/tested/handlers"

func TestHandlersTest(t *testing.T) {
	assert.Expect(t, true, true)
}

func TestHandlers(t *testing.T) {

	tests := []struct {
		handler http.Handler
		exp     int
	}{
		{handlers.OkHandler, http.StatusOK},
		{handlers.NotFoundHandler, http.StatusNotFound},
		{handlers.StatusCodeHandler(http.StatusOK), http.StatusOK},
	}

	for _, test := range tests {

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)

		test.handler.ServeHTTP(w, req)
		resp := w.Result()

		assert.Expect(t, test.exp, resp.StatusCode)
	}
}
