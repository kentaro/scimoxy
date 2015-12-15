package scimoxy

import (
	"reflect"
	"testing"
)

func typeOf(s interface{}) string {
	return reflect.TypeOf(s).String()
}

func testNewServer(t *testing.T) {
	actual := typeOf(newServer())
	expected := "scimoxy.Server"

	if actual != expected {
		t.Fatalf("type mismatched: %v", actual)
	}
}

func testServiceHandler(t *testing.T) {
	actual := typeOf(newServer().serviceHandlerFor("/slack/Users"))
	expected := "scimoxy.SlackHandler"

	if actual != expected {
		t.Fatalf("type mismatched: %v", actual)
	}
}
