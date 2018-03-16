package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePage(t *testing.T) {
	handler := http.HandlerFunc(HomePage)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "http://example.com", nil)

	handler.ServeHTTP(w, r)

	if got, want := w.Code, 200; got != want {
		t.Errorf("Code got %d; want %d", got, want)
	}

	if got, want := w.Body.String(), "Hi, everyone"; got != want {
		t.Errorf("Body got %s; want %s", got, want)
	}
}
