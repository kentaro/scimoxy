package scimoxy

import (
	"testing"
)

func TestNewServer(t *testing.T) {
	actual := typeOf(NewServer())
	expected := "*scimoxy.Server"

	if actual != expected {
		t.Fatalf("type mismatched: %v", actual)
	}
}
