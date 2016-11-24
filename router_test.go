package aero

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//Test wrong route
func TestRoute(t *testing.T) {
	testRoute := Router()
	routeIsValid := true
	testRoute.Get("/bingo/world", func(res http.ResponseWriter, req *http.Request) {
		routeIsValid = false
	})

	req, err := http.NewRequest("GET", "/hello/bingo", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	testRoute.ServeHTTP(res, req)

	if !routeIsValid {
		t.Error("Invalid route")
	}
}

//Test define path with trailing slash and request without trailing slash
func TestPathTrailingSlash(t *testing.T) {
	testRoute := Router()
	trailingSlashError := true
	testRoute.Get("/aero/route/", func(res http.ResponseWriter, req *http.Request) {
		trailingSlashError = false
	})

	req, err := http.NewRequest("GET", "/aero/route", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	testRoute.ServeHTTP(res, req)

	if trailingSlashError {
		t.Error("Trailing slash matters")
	}
}

type muxTest struct {
	path        string
	isPathMatch bool
}

func TestMuxMatchingPath(t *testing.T) {
	testRouter := Router()
	testHandler := func(w http.ResponseWriter, r *http.Request) {}
	simulateRequest := testRouter.registerHandler("GET", "/bingo", testHandler)

	testReq := []muxTest{
		{"/bingo", true},
		{"/bingo/", true},
		{"/binggo", false},
	}

	for _, test := range testReq {
		match, _ := simulateRequest.matchingPath(test.path)
		if match != test.isPathMatch {
			t.Error("Fail at", test.path, "Expected match is", test.isPathMatch, "instead of", match)
		}
	}
}
