package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestGetUserEndpoint(t *testing.T) {

// 	req, err := http.NewRequest("GET", "/user", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	q := req.URL.Query()
// 	// q.Add("{i", "61619afc64760c13a706290e")
// 	q= path.Join(q.Path, "61619afc64760c13a706290e")

// 	req.URL.RawQuery = q.Encode()
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetUserEndpoint)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	expected := `[{"_id": "61619afc64760c13a706290e","name": "Sumit","email": "sumit0311@gmail.com","password": "$2a$14$7dYCuoirp2ZF/WaJE56rsOvYl1A3RDqxcS7SUq3v5r4a/hZk6qNuy"}]`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }
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