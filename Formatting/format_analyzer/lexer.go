package format_analyzer

import (
	"fmt"
	"io"
	"slices"
	"strings"

	gcch "github.com/PlayerR9/go-commons/runes"
)

type Lexer struct {
	prefix        rune
	allowed_verbs []rune
	chars         []rune
}

func (l *Lexer) next_rune() (rune, bool) {
	if len(l.chars) == 0 {
		return 0, false
	}

	c := l.chars[0]
	l.chars = l.chars[1:]

	return c, true
}

func (l *Lexer) peek_rune() (rune, bool) {
	if len(l.chars) == 0 {
		return 0, false
	}

	return l.chars[0], true
}

func (l *Lexer) lex_one() (*Token, error) {
	c, ok := l.next_rune()
	if !ok {
		return nil, io.EOF
	}

	if c != l.prefix {
		var builder strings.Builder

		builder.WriteRune(c)

		for {
			c, ok := l.peek_rune()
			if !ok || c == l.prefix {
				break
			}

			builder.WriteRune(c)
			_, _ = l.next_rune()
		}

		tk := NewToken(false, builder.String())

		return tk, nil
	}

	next_c, ok := l.next_rune()
	if !ok || next_c == l.prefix {
		tk := NewToken(false, string(l.prefix))

		return tk, nil
	}

	var tk *Token

	_, ok = slices.BinarySearch(l.allowed_verbs, next_c)
	if !ok {
		return nil, fmt.Errorf("flag \"%c%c\" is not supported", c, next_c)
	}

	tk = NewToken(true, string(next_c))

	return tk, nil
}

func (l *Lexer) lex() ([]*Token, error) {
	var tokens []*Token

	for {
		tk, err := l.lex_one()
		if err == io.EOF {
			break
		} else if err != nil {
			return tokens, err
		}

		if tk == nil {
			continue
		}

		if tk.IsVerb || len(tokens) == 0 {
			tokens = append(tokens, tk)
		}

		prev := tokens[len(tokens)-1]
		if prev.IsVerb {
			tokens = append(tokens, tk)
		} else {
			prev.Data += tk.Data
		}
	}

	return tokens, nil
}

func (l *Lexer) Format(format string, data Formatter) (string, error) {
	chars, err := gcch.StringToUtf8(format)
	if err != nil {
		return "", fmt.Errorf("invalid format: %w", err)
	}

	l.chars = chars

	tokens, err := l.lex()
	if err != nil {
		return "", fmt.Errorf("invalid format: %w", err)
	}

	res, err := apply(tokens, data)
	if err != nil {
		return "", err
	}

	return res, nil
}
