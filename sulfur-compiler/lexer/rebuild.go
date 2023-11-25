package lexer

import "strings"

// this is a diagnostic tool. It "reconstructs" the source code from the lexed tokens.

func RebuildFromTokens(tokens []Token) string {
	orig := ""
	startOfLine := true
	binding := false
	for _, t := range tokens {
		switch t.tokenType {
		case TT_STANDING_SYMBOL:
			if !startOfLine {
				orig += " "
			}
			orig += t.content
			binding = false
		case TT_OPEN_SYMBOL:
			if !startOfLine && !binding {
				orig += " "
			}
			orig += t.content
			binding = false
		case TT_CLOSE_SYMBOL:
			if !startOfLine {
				orig += " "
			}
			orig += t.content
			binding = false
		case TT_OPEN_BIND_SYMBOL:
			orig += t.content
			binding = false
		case TT_BINDING_SYMBOL:
			orig += t.content
			binding = true
		case TT_IDENT:
			if !startOfLine && !binding {
				orig += " "
			}
			orig += t.content
			binding = false
		case TT_STR_LIT:
			if !startOfLine && !binding {
				orig += " "
			}
			orig += "\"" + t.content + "\""
			binding = false
		case TT_NUMSTR_LIT:
			if !startOfLine && !binding {
				orig += " "
			}
			orig += t.content
			binding = false
		case TT_SYNTAX_ERROR:
			orig += "\n\nERROR:\n\n" + t.content
		case TT_COMMENT:
			// should not happen
		case TT_WHITESPACE:
			// should not happen
		case TT_INDENT_LINE:
			orig += "\n" + strings.Repeat("  ", t.indent)
			binding = false
		}
		if t.tokenType != TT_INDENT_LINE {
			startOfLine = false
		} else {
			startOfLine = true
		}
	}
	return orig
}
