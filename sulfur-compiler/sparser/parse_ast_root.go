package sparser

import (
	"sulfur-compiler/lexer"
)

func parseAstRoot(cursor *parseCursor) error {
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		return parseAstStatementStartChild(cursor)
	//case lexer.TT_OPEN_SYMBOL:
	//case lexer.TT_CLOSE_SYMBOL:
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	case lexer.TT_IDENT:
		return parseAstStatementStartChild(cursor)
		//case lexer.TT_STR_LIT:
		//case lexer.TT_NUMSTR_LIT:
		//case lexer.TT_SYNTAX_ERROR:
		//case lexer.TT_COMMENT:
		//case lexer.TT_WHITESPACE:
		//case lexer.TT_LINE:
	}
	return parseError(cursor, "MAIN", "unhandled parse")
}
