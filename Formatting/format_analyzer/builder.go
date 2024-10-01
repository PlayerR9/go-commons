package format_analyzer

import "slices"

const (
	DefaultPrefix rune = '%'
)

type Builder struct {
	prefix        rune
	allowed_verbs []rune
}

func (b *Builder) SetPrefix(char rune) {
	b.prefix = char
}

func (b *Builder) Register(verb rune) {
	pos, ok := slices.BinarySearch(b.allowed_verbs, verb)
	if ok {
		return
	}

	b.allowed_verbs = slices.Insert(b.allowed_verbs, pos, verb)
}

func (b Builder) Build() *Lexer {
	var prefix rune

	if prefix == 0 {
		prefix = DefaultPrefix
	}

	return &Lexer{
		prefix: prefix,
	}
}

func (b *Builder) Reset() {
	b.prefix = 0
}
