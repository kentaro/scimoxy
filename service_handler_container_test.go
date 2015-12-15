package scimoxy

import (
	"testing"
)

func TestServiceHandlerContainerAdd(t *testing.T) {
	c := NewServiceHandlerContainer()

	{
		actual := c.Count()
		expected := 1

		if actual != expected {
			t.Fatalf("expected %v but got %v", expected, actual)
		}
	}

	{
		c.Add("__test__", NewSlackHandler())
		actual := c.Count()
		expected := 2

		if actual != expected {
			t.Fatalf("expected %v but got %v", expected, actual)
		}
	}
}

func TestServiceHandlerContainerGet(t *testing.T) {
	c := NewServiceHandlerContainer()
	h := NewSlackHandler()
	c.Add("__test__", h)

	actual := c.Get("__test__")
	expected := h

	if actual != expected {
		t.Fatalf("expected %v but got %v", expected, actual)
	}
}
