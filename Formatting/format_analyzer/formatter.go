package format_analyzer

import (
	"fmt"
	"strings"
)

type Formatter interface {
	Format(verb string) (string, error)
}

func apply(tokens []*Token, data Formatter) (string, error) {
	var builder strings.Builder

	if data == nil {
		for _, token := range tokens {
			if token.IsVerb {
				return builder.String(), fmt.Errorf("flag \"%s\" is not supported", token.Data)
			}

			builder.WriteString(token.Data)
		}
	} else {
		for _, token := range tokens {
			if !token.IsVerb {
				builder.WriteString(token.Data)
				continue
			}

			res, err := data.Format(token.Data)
			if err != nil {
				return builder.String(), err
			}

			builder.WriteString(res)
		}
	}

	return builder.String(), nil
}
