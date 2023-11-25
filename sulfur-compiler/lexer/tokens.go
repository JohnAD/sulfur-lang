package lexer

import "fmt"

type TokenType int

const (
	TT_STANDING_SYMBOL TokenType = iota
	TT_OPEN_SYMBOL
	TT_CLOSE_SYMBOL
	TT_OPEN_BIND_SYMBOL
	TT_BINDING_SYMBOL
	TT_IDENT
	TT_STR_LIT
	TT_NUMSTR_LIT
	TT_SYNTAX_ERROR
	TT_COMMENT
	TT_WHITESPACE
	TT_INDENT_LINE
)

func (tt TokenType) String() string {
	switch tt {
	case TT_STANDING_SYMBOL:
		return "STANDING_SYMBOL"
	case TT_OPEN_SYMBOL:
		return "OPEN_SYMBOL"
	case TT_CLOSE_SYMBOL:
		return "CLOSE_SYMBOL"
	case TT_OPEN_BIND_SYMBOL: // such as a(
		return "OPEN_BIND_SYMBOL"
	case TT_BINDING_SYMBOL: // such as a::b or a\b
		return "BINDING_SYMBOL"
	case TT_IDENT:
		return "IDENT"
	case TT_STR_LIT:
		return "STR_LIT"
	case TT_NUMSTR_LIT:
		return "NUMSTR_LIT"
	case TT_SYNTAX_ERROR:
		return "SYNTAX_ERROR"
	case TT_COMMENT:
		return "COMMENT"
	case TT_WHITESPACE:
		return "WHITESPACE"
	case TT_INDENT_LINE:
		return "INDENT_LINE"
	default:
		return fmt.Sprintf("%d", int(tt))
	}
}

type Token struct {
	tokenType TokenType
	//sourceFile string TODO: later this needs to be encoded
	//sourceLine int
	//sourceOffset int
	content string
	indent  int
}

func NewToken(init TokenType) Token {
	t := Token{tokenType: init, content: ""}
	return t
}

func NewTokenWithRune(init TokenType, ch rune) Token {
	t := Token{tokenType: init, content: string(ch)}
	return t
}

func (t Token) String() string {
	if t.tokenType == TT_INDENT_LINE {
		return fmt.Sprintf("{ %v==%d }", t.tokenType, t.indent)
	}
	return fmt.Sprintf("{ %v \"%s\" }", t.tokenType, t.content)
}
