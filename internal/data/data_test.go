package data_test

import (
	"testing"

	"github.com/snaztoz/granary/internal/data"
)

func TestDataRepresentation(t *testing.T) {
	d := data.T{
		"foo": "bar",
		"baz": "bat",
	}

	if d.String() != "baz\nfoo" {
		t.Fatalf("representation mismatched, get '%s' instead", d)
	}
}
