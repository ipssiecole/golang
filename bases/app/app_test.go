package app_test

import (
	"bases/app"
	"testing"
)

// go test -v ./...
func TestSayHello(t *testing.T) {
	got := app.SayHello("John")
	want := "Hello John"

	if got != want {
		t.Errorf("SayHello() got %s; want %s", got, want)
	}
}
