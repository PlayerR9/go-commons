package format_analyzer

type Token struct {
	IsVerb bool
	Data   string
}

func NewToken(isVerb bool, data string) *Token {
	return &Token{
		IsVerb: isVerb,
		Data:   data,
	}
}
