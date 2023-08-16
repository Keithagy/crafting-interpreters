package lex

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any // TODO: Looks like a good place for a generic
	Line    int
}

func (t Token) String() string {
	return t.Type.String() + " " + t.Lexeme + " " + t.Literal.(string)
}
