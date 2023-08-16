package lex

type Scanner struct {
	source string
	tokens []Token

	// offsets that index into source.
	// start points to the first character in the current lexeme
	// current points at character currently being considered
	// REMEMBER: a lexeme can be more than 1 character, such as in the case of a 3-digit integer literal
	start, current int
	line           int //
}

func New(source string) Scanner {
	return Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}
func (s *Scanner) scanTokens() []Token {
	for i, c := range s.source {
		s.start = i
		s.current = i
		s.scanToken(c)
	}

	s.tokens = append(s.tokens, Token{Eof, "", nil, s.line})
	return s.tokens
}

func (s *Scanner) scanToken(c rune) {
	switch c {
	// TODO: clear opportunity for cleaner mappings between token types and lexemes
	case '(':
		s.addTokenWithoutLiteral(LeftParen, string(c))
	case ')':
		s.addTokenWithoutLiteral(RightParen, string(c))
	case '{':
		s.addTokenWithoutLiteral(LeftBrace, string(c))
	case '}':
		s.addTokenWithoutLiteral(RightBrace, string(c))
	case ',':
		s.addTokenWithoutLiteral(Comma, string(c))
	case '.':
		s.addTokenWithoutLiteral(Dot, string(c))
	case '-':
		s.addTokenWithoutLiteral(Minus, string(c))
	case '+':
		s.addTokenWithoutLiteral(Plus, string(c))
	case ';':
		s.addTokenWithoutLiteral(Semicolon, string(c))
	case '*':
		s.addTokenWithoutLiteral(Star, string(c))
	}
}

func (s *Scanner) addTokenWithoutLiteral(t TokenType, lexeme string) {
	s.addToken(t, lexeme, nil)
}

func (s *Scanner) addToken(t TokenType, lexeme string, literal any) {
	s.tokens = append(s.tokens, Token{t, lexeme, literal, s.line})
}
