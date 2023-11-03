package scanner

import "fmt"

// Token represents a lexical token.
type Token struct {
	Literal interface{}
	Lexeme  string
	Type    TokenType
	Line    int
}

// TokenType is an enumeration of all the possible types of tokens.
type TokenType int

// Enumerated values for TokenType.
const (
	// Define token types here, for example:
	TokenEOF TokenType = iota
	TokenIdentifier
	// ... and so on for each token type.
)

// String representation of the Token for printing.
func (t Token) String() string {
	return fmt.Sprintf("%d %v %v", t.Line, t.Type, t.Lexeme)
}

// Scanner represents a lexical scanner.
type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

// New creates and returns a new Scanner.
func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

// ScanTokens breaks the source into tokens and returns them.
func (s *Scanner) ScanTokens() []Token {
	// Implement the tokenization logic here.
	// This is just a placeholder to illustrate the idea.
	for s.current < len(s.source) {
		// Scan a token and add to s.tokens
	}

	// Add an end-of-file token at the end
	s.tokens = append(s.tokens, Token{Type: TokenEOF, Line: s.line})
	return s.tokens
}
