package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolneItemValue(cursor *parseCursor, token lexer.Token) error {
	debug(AST_ROLNE_ITEM_VALUE, "MAIN", cursor)
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if token.Content == "," {
			return childParseAstRolneItemFinish(cursor, token)
		}
	case lexer.TT_BINDING_SYMBOL:
		return binderHandlingInPlace(cursor, token)
	case lexer.TT_CLOSE_SYMBOL:
		return childParseAstRolneItemFinish(cursor, token) // a closing "}" etc found
	case lexer.TT_IDENT:
		return parseAstRolneItemNameAssignValue(cursor, token, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemNameAssignValue(cursor, token, ASTN_STRLIT)
	case lexer.TT_LINE:
		return childParseAstRolneItemFinish(cursor, token)
	}
	return fmt.Errorf("unhandled AST_ROLNE_ITEM_VALUE parse of %v", token)
}

func parseAstRolneItemValueStartViaEqualSign(cursor *parseCursor) error {
	// when this is called, we should be currently pointing to the AST_ROLNE_ITEM_NAME or AST_ROLNE_ITEM_TYPE
	debug(AST_ROLNE_ITEM_VALUE, "PARIVSVES", cursor)
	_ = finishAstNode(cursor) // finish pointing to AST_ROLNE_ITEM_*
	return gotoChild(cursor, ROLEITEM_VALUECHILD)
}

func parseAstRolneItemValueStartViaBinding(cursor *parseCursor, token lexer.Token) error {
	// this should only be called from AST_ROLNE_ITEM_NAME
	debug(AST_ROLNE_ITEM_VALUE, "PARIVSVB", cursor)
	_ = finishAstNode(cursor) // finish pointing to AST_ROLNE_ITEM_*
	parseAstRolneItemMoveNameToChild(cursor)
	_ = gotoChild(cursor, ROLEITEM_VALUECHILD) // point to AST_ROLNE_ITEM_VALUE
	return binderHandlingInPlace(cursor, token)
}

func parseAstRolneItemNameAssignValue(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	debug(AST_ROLNE_ITEM_VALUE, "PARINAV", cursor)
	if cursor.currentNode.Nature != ASTN_NULL {
		return childParseAstRolneItemFinish(cursor, token)
	}
	cursor.currentNode.Name = token.Content
	cursor.currentNode.src = token
	cursor.currentNode.Nature = nature
	return nil
}
