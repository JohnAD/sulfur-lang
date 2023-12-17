package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstOrderedBindingChild(cursor *parseCursor) error {
	debug(AST_ORDERED_BINDING_CHILD, "MAIN", cursor)
	switch cursor.src.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//case lexer.TT_OPEN_SYMBOL:
	//case lexer.TT_CLOSE_SYMBOL:
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	case lexer.TT_IDENT:
		return finishAstOrderedBindingChild(cursor, cursor.src, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return finishAstOrderedBindingChild(cursor, cursor.src, ASTN_STRLIT)
		//case lexer.TT_NUMSTR_LIT:
		//case lexer.TT_SYNTAX_ERROR:
		//case lexer.TT_COMMENT:
		//case lexer.TT_WHITESPACE:
		//case lexer.TT_LINE:
	}
	return fmt.Errorf("unhandled AST_ORDRED_BINDING_CHILD parse of %v", cursor.src)
}

func finishAstOrderedBindingChild(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	cursor.currentNode.Name = token.Content
	cursor.currentNode.src = token
	cursor.currentNode.Nature = nature
	return finishAstNode(cursor)
}
