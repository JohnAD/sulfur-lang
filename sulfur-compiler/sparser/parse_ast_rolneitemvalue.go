package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolneItemValue(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if token.Content == "," {
			return childParseAstRolneItemFinish(cursor, token)
		}
	case lexer.TT_CLOSE_SYMBOL:
		return childParseAstRolneItemFinish(cursor, token) // a closing "}" etc found
	case lexer.TT_IDENT:
		return parseAstRolneItemNameAssignValue(cursor, token, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemNameAssignValue(cursor, token, ASTN_STRLIT)
	case lexer.TT_LINE:
		return childParseAstRolneItemFinish(cursor, token)
	}
	return fmt.Errorf("unhandled AST_ROLNE_VALUE parse of %v", token)
}

func parseAstRolneItemValueStartViaEqualSign(cursor *parseCursor) error {
	// when this is called, we should be currently pointing to the AST_ROLNE_NAME or AST_ROLNE_TYPE
	_ = finishAstNode(cursor) // finish pointing to AST_ROLNE_ITEM
	return gotoChild(cursor, ROLEITEM_VALUECHILD)
}

func parseAstRolneItemNameAssignValue(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	debug("PARINAV-start", cursor)
	if cursor.currentNode.Nature != ASTN_NULL {
		return childParseAstRolneItemFinish(cursor, token)
	}
	cursor.currentNode.Name = token.Content
	cursor.currentNode.Nature = nature
	return nil
}
