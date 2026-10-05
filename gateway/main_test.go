package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	for _, method := range []string{http.MethodGet, http.MethodPost} {
		t.Run(method, func(t *testing.T) {
			response := httptest.NewRecorder()
			pingHandler(response, httptest.NewRequest(method, "/ping", nil))
			if method == http.MethodGet {
				if response.Code != http.StatusOK || response.Body.String() != "pong" {
					t.Fatalf("got status=%d body=%q; want 200 and pong", response.Code, response.Body.String())
				}
			} else if response.Code != http.StatusMethodNotAllowed || response.Header().Get("Allow") != "GET" {
				t.Fatalf("got status=%d Allow=%q; want 405 and GET", response.Code, response.Header().Get("Allow"))
			}
		})
	}
}
