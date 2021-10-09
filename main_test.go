package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestGetUsers(t *testing.T) {
    t.Run("returns Pepper's score", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/user/61619afc64760c13a706290e", nil)
        response := httptest.NewRecorder()

        GetUserEndpoint(response, request)

        got := response.Body.String()
        want := `{
			"_id": "61619afc64760c13a706290e",
			"name": "Sumit",
			"email": "sumit0311@gmail.com",
			"password": "$2a$14$7dYCuoirp2ZF/WaJE56rsOvYl1A3RDqxcS7SUq3v5r4a/hZk6qNuy"
		}`

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })
}