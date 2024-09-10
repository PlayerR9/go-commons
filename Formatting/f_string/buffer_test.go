package f_string

import (
	"fmt"
	"strings"
	"testing"
)

type MockElement struct {
}

func (m *MockElement) FString(trav *Traversor, opts ...Option) error {
	err := trav.AddLine("MockElement")
	if err != nil {
		return err
	}

	return nil
}

func TestBufferIndent(t *testing.T) {
	p, trav := NewStdPrinter(
		NewFormatter(NewIndentConfig("   ", 1)),
	)

	err := trav.AddJoinedLine(" ", "\t", "a", "\t", "b", "c")
	if err != nil {
		t.Fatalf("Expected no error, but got %s", err.Error())
	}

	form, _ := trav.GetConfig(
		WithModifiedIndent(2),
	)

	err = ApplyForm(
		form,
		trav,
		&MockElement{},
	)
	if err != nil {
		t.Fatalf("Expected no error, but got %s", err.Error())
	}

	pages := Stringfy(p.GetPages(), 1)

	fmt.Println(strings.Join(pages, "\n"))

	t.Fatalf("Test not implemented")
}
