package internal

import "strings"

type StackTrace struct {
	Trace []string
}

func (st StackTrace) String() string {
	return strings.Join(st.Trace, " -> ")
}
