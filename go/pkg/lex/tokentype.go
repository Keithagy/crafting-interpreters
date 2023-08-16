package lex

type TokenType uint

const (
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star
	Bang
	BangEq
	Eq
	EqEq
	Greater
	GreaterEq
	Less
	LessEq
	Identifier
	String
	Number
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While
	Eof
)

func (t TokenType) String() string {
	switch t {
	case LeftParen:
		return "LeftParen"
	case RightParen:
		return "RightParen"
	case LeftBrace:
		return "LeftBrace"
	case RightBrace:
		return "RightBrace"
	case Comma:
		return "Comma"
	case Dot:
		return "Dot"
	case Minus:
		return "Minus"
	case Plus:
		return "Plus"
	case Semicolon:
		return "Semicolon"
	case Slash:
		return "Slash"
	case Star:
		return "Star"
	case Bang:
		return "Bang"
	case BangEq:
		return "BangEq"
	case Eq:
		return "Eq"
	case EqEq:
		return "EqEq"
	case Greater:
		return "Greater"
	case GreaterEq:
		return "GreaterEq"
	case Less:
		return "Less"
	case LessEq:
		return "LessEq"
	case Identifier:
		return "Identifier"
	case String:
		return "String"
	case Number:
		return "Number"
	case And:
		return "And"
	case Class:
		return "Class"
	case Else:
		return "Else"
	case False:
		return "False"
	case Fun:
		return "Fun"
	case For:
		return "For"
	case If:
		return "If"
	case Nil:
		return "Nil"
	case Or:
		return "Or"
	case Print:
		return "Print"
	case Return:
		return "Return"
	case Super:
		return "Super"
	case This:
		return "This"
	case True:
		return "True"
	case Var:
		return "Var"
	case While:
		return "While"
	case Eof:
		return "Eof"
	default:
		return "Unknown"
	}
}

func (t TokenType) ScanTarget() string {
	switch t {
	case LeftParen:
		return "("
	case RightParen:
		return ")"
	case LeftBrace:
		return "{"
	case RightBrace:
		return "}"
	case Comma:
		return ","
	case Dot:
		return "."
	case Minus:
		return "-"
	case Plus:
		return "+"
	case Semicolon:
		return ";"
	case Slash:
		return "/"
	case Star:
		return "*"
	case Bang:
		return "!"
	case BangEq:
		return "!="
	case Eq:
		return "="
	case EqEq:
		return "=="
	case Greater:
		return ">"
	case GreaterEq:
		return ">="
	case Less:
		return "<"
	case LessEq:
		return "<="
	case Identifier:
		return "var"
	case String:
		return "str"
	case Number:
		return "num"
	case And:
		return "and"
	case Class:
		return "Class"
	case Else:
		return "else"
	case False:
		return "false"
	case Fun:
		return "fun"
	case For:
		return "for"
	case If:
		return "if"
	case Nil:
		return "nil"
	case Or:
		return "or"
	case Print:
		return "print"
	case Return:
		return "return"
	case Super:
		return "super"
	case This:
		return "this"
	case True:
		return "true"
	case Var:
		return "var"
	case While:
		return "while"
	case Eof:
		return "EOF"
	default:
		return "Unknown"
	}
}
