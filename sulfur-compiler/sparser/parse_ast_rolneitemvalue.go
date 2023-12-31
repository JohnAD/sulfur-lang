package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolneItemValue(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if cursor.src.Content == "," {
			return childParseAstRolneItemFinish(cursor)
		}
	case lexer.TT_BINDING_SYMBOL:
		return binderHandlingInPlace(cursor)
	case lexer.TT_CLOSE_SYMBOL:
		return childParseAstRolneItemFinish(cursor) // a closing "}" etc found
	case lexer.TT_IDENT:
		return parseAstRolneItemNameAssignValue(cursor, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemNameAssignValue(cursor, ASTN_STRLIT)
	case lexer.TT_LINE:
		return childParseAstRolneItemFinish(cursor)
	}
	return fmt.Errorf("unhandled AST_ROLNE_ITEM_VALUE parse of %v", cursor.src)
}

func parseAstRolneItemValueStartViaEqualSign(cursor *parseCursor) error {
	// when this is called, we should be currently pointing to the AST_ROLNE_ITEM_NAME or AST_ROLNE_ITEM_TYPE
	debug(cursor, "PARIVSVES")
	_ = finishAstNode(cursor) // finish pointing to AST_ROLNE_ITEM_*
	return gotoChild(cursor, ROLEITEM_VALUECHILD)
}

func parseAstRolneItemValueStartViaBinding(cursor *parseCursor) error {
	// this should only be called from AST_ROLNE_ITEM_NAME
	debug(cursor, "PARIVSVB")
	_ = finishAstNode(cursor) // finish pointing to AST_ROLNE_ITEM_*
	parseAstRolneItemMoveNameToChild(cursor)
	_ = gotoChild(cursor, ROLEITEM_VALUECHILD) // point to AST_ROLNE_ITEM_VALUE
	return binderHandlingInPlace(cursor)
}

func parseAstRolneItemNameAssignValue(cursor *parseCursor, nature AstNodeNature) error {
	debug(cursor, "PARINAV")
	if cursor.currentNode.Nature != ASTN_NULL {
		return childParseAstRolneItemFinish(cursor)
	}
	cursor.currentNode.Src = cursor.src
	cursor.currentNode.Nature = nature
	return nil
}
