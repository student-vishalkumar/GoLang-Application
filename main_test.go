package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPageHandlers(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		handler http.HandlerFunc
	}{
		{"home page", "/home", homePage},
		{"courses page", "/courses", coursePage},
		{"about page", "/about", aboutPage},
		{"contact page", "/contact", contactPage},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			tc.handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("%s: got status %v want %v", tc.path, status, http.StatusOK)
			}

			expected := "text/html; charset=utf-8"
			if contentType := rr.Header().Get("Content-Type"); contentType != expected {
				t.Errorf("%s: got content type %v want %v", tc.path, contentType, expected)
			}

			if rr.Body.Len() == 0 {
				t.Errorf("%s: response body is empty", tc.path)
			}
		})
	}
}