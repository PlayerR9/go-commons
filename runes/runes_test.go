package runes

import (
	"slices"
	"testing"
)

func TestEitherOrString(t *testing.T) {
	var (
		Input []rune = []rune{'a', 'b', 'c'}
	)

	const (
		Expected string = "either a, b, or c"
	)

	res := EitherOrString(Input, false)
	if res != Expected {
		t.Fatalf("EitherOrString(%q) = %q; want %q", Input, res, Expected)
	}
}

func TestNormalizeRunes(t *testing.T) {
	var (
		Input    []rune = []rune{'a', 'b', 'c', '\r', '\n', 'b', '\t'}
		Expected []rune = []rune{'a', 'b', 'c', '\n', 'b', '\t'}
	)

	res, err := NormalizeRunes(Input)
	if err != nil {
		t.Fatalf("NormalizeRunes(%q) = %q; want %q", Input, res, Expected)
	}

	if !slices.Equal(res, Expected) {
		t.Fatalf("NormalizeRunes(%q) = %q; want %q", Input, res, Expected)
	}
}
