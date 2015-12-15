package scimoxy

import (
	"reflect"
	"testing"
)

func testNewServer(t *testing.T) {
	actual := reflect.TypeOf(newServer()).String()
	expected := "scimoxy.Server"

	if actual != expected {
		t.Fatalf("type mismatched: %v", actual)
	}
}
