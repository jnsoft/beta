package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	str := "This is a string"

	reversed_str := Reverse(str)

	reversed_reversed_str := Reverse(reversed_str)

	if str == reversed_str {
		t.Fatalf("Reversed string match original. Got %s, want reverse of %s", reversed_str, str)
	}

	if str != reversed_reversed_str {
		t.Fatalf("Reversed reversed string does not match original. Got %s, want %s", reversed_reversed_str, str)
	}
}
