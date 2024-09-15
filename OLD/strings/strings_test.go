package strings

import (
	"testing"

	gcstr "github.com/PlayerR9/go-commons/strings"
)

func TestFindContentIndexes(t *testing.T) {
	const (
		OpToken string = "("
		ClToken string = ")"
	)

	var (
		ContentTokens []string = []string{
			"(", "(", "a", "+", "b", ")", "*", "c", ")", "+", "d",
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

func TestOrString(t *testing.T) {
	TestValues := []string{"a", "b", "c "}

	str := gcstr.OrString(TestValues, false)
	if str != "a, b, or c" {
		t.Errorf("OrString(%q) = %q; want %q", TestValues, str, "a, b, or c")
	}
}

func TestAdaptToScreeenWidth(t *testing.T) {
	adapted, cut := AdaptToScreenWidth([]string{"a", "bb", "ccc", "dddd"}, 13, " ")
	if cut != 1 {
		t.Errorf("expected 1, got %d instead", cut)
	}

	if adapted != "a bb ... dddd" {
		t.Errorf("expected %q, got %q instead", "a bb ... dddd", adapted)
	}
}
