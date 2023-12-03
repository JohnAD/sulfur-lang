package lexer

import "unicode"

type RuneCategory int

const (
	RC_WHITESPACE = iota
	RC_LETTER
	RC_NUMBER
	RC_PUNCTUATION
	RC_OPEN_PUNCTUATION
	RC_CLOSE_PUNCTUATION
	RC_LINE_END
	RC_QUOTE
	RC_FORBIDDEN
)

func GetRuneCategory(ch rune) RuneCategory {
	if ch == '\n' {
		return RC_LINE_END
	}
	if ch == '"' {
		return RC_QUOTE
	}
	if ch == '(' || ch == '[' || ch == '{' {
		return RC_OPEN_PUNCTUATION
	}
	if ch == ')' || ch == ']' || ch == '}' {
		return RC_CLOSE_PUNCTUATION
	}
	if ch == '_' {
		return RC_LETTER
	}
	if unicode.IsSpace(ch) {
		return RC_WHITESPACE
	}
	if unicode.IsLetter(ch) {
		return RC_LETTER
	}
	if unicode.IsNumber(ch) {
		return RC_NUMBER
	}
	if unicode.IsPunct(ch) {
		return RC_PUNCTUATION
	}
	if unicode.IsSymbol(ch) {
		return RC_PUNCTUATION
	}
	return RC_FORBIDDEN
}

// TODO: to add:
//   comment prefix
//   closing punctuation
//
