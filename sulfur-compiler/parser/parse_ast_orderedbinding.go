package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstOrderedBinding(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//case lexer.TT_OPEN_SYMBOL:
	//case lexer.TT_CLOSE_SYMBOL:
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	case lexer.TT_IDENT:
		return bindSecondHalfAndFinish(cursor, token)
		//case lexer.TT_STR_LIT:
		//case lexer.TT_NUMSTR_LIT:
		//case lexer.TT_SYNTAX_ERROR:
		//case lexer.TT_COMMENT:
		//case lexer.TT_WHITESPACE:
		//case lexer.TT_INDENT_LINE:
	}
	return fmt.Errorf("unhandled AST_BINDING parse of %v", token)
}

func bindSecondHalfAndFinish(cursor *parseCursor, token lexer.Token) error {
	addChild(cursor, AST_IDENTIFIER, ASTN_IDENTIFIER, token.Content)
	return finishAstNode(cursor)
}

func binderHandling(cursor *parseCursor, token lexer.Token) error {
	return becomeLastChildMakePreviousChildAChildThenBecomeChild(cursor, AST_ORDERED_BINDING, ASTN_META_BINDING, token.Content)
}

func binderHandlingInPlace(cursor *parseCursor, token lexer.Token) error {
	return swapSelfMakingPreviousAChild(cursor, AST_ORDERED_BINDING, ASTN_META_BINDING, token.Content)
}
