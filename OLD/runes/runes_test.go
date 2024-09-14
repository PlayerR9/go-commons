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

func TestFindContentIndexes(t *testing.T) {
	const (
		OpToken rune = '('
		ClToken rune = ')'
	)

	var (
		ContentTokens []rune = []rune{
			'(', '(', 'a', '+', 'b', ')', '*', 'c', ')', '+', 'd',
		}
	)

	indices, err := FindContentIndexes(OpToken, ClToken, ContentTokens)
	if err != nil {
		t.Errorf("expected no error, got %s instead", err.Error())
	}

	if indices[0] != 1 {
		t.Errorf("expected 1, got %d instead", indices[0])
	}

	if indices[1] != 9 {
		t.Errorf("expected 9, got %d instead", indices[1])
	}
}
