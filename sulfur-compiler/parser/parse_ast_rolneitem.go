package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolneItem(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//case lexer.TT_OPEN_SYMBOL:
	//case lexer.TT_CLOSE_SYMBOL:
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	case lexer.TT_IDENT:
		return parseAstRolneItemValue(cursor, token)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil
	case lexer.TT_INDENT_LINE:
		return finishAstRoleItem(cursor)
	}
	return fmt.Errorf("unhandled AST_ROLNE_ITEM parse of %v", token)
}

func parseAstRolneItemValue(cursor *parseCursor, token lexer.Token) error {
	addChild(cursor, AST_IDENTIFIER, ASTN_IDENTIFIER, token.Content)
	return nil
}

func finishAstRoleItem(cursor *parseCursor) error {
	return finishAstNode(cursor)
}

func parseAstRolneItemStartAssignment(cursor *parseCursor, token lexer.Token) error {
	return becomeLastChildMakePreviousChildAChildThenBecomeChild(cursor, AST_ROLNE_ITEM, ASTN_INFIX_OPERATOR, token.Content)
}
