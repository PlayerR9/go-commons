package runes

import (
	"testing"
)

func TestToInt(t *testing.T) {
	digit, ok := ToInt('A')
	if !ok {
		t.Error("should convert 'A' to an integer")
	}

	if digit != 10 {
		t.Errorf("expected 10 but got %d", digit)
	}

	char, ok := FromInt(digit)
	if !ok {
		t.Error("should convert 10 to a rune")
	}

	if char != 'a' {
		t.Errorf("expected 'a' but got %c", char)
	}
}
