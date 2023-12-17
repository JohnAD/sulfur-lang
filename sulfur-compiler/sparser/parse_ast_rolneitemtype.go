package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolneItemType(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if cursor.src.Content == "=" {
			return parseAstRolneItemValueStartViaEqualSign(cursor)
		} else if cursor.src.Content == "," {
			return childParseAstRolneItemFinish(cursor)
		}
	case lexer.TT_CLOSE_SYMBOL:
		return childParseAstRolneItemFinish(cursor) // a closing "}" etc. found
	case lexer.TT_IDENT:
		return parseAstRolneItemTypeAssignType(cursor, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemTypeAssignType(cursor, ASTN_STRLIT)
	case lexer.TT_LINE:
		return childParseAstRolneItemFinish(cursor)
	}
	return fmt.Errorf("unhandled AST_ROLNE_ITEM_VALUE parse of %v", cursor.src)
}

func parseAstRolneItemTypeStart(cursor *parseCursor) error {
	// when this is called, we should be currently pointing to the AST_ROLNE_ITEM_NAME
	debug(cursor, "PARITS")
	_ = finishAstNode(cursor)                    // finish and point to AST_ROLNE_ITEM
	return gotoChild(cursor, ROLEITEM_TYPECHILD) // now point to AST_ROLNE_ITEM_TYPE
}

func parseAstRolneItemTypeAssignType(cursor *parseCursor, nature AstNodeNature) error {
	debug(cursor, "PARITAV")
	if cursor.currentNode.Nature != ASTN_NULL {
		return childParseAstRolneItemFinish(cursor)
	}
	cursor.currentNode.Name = cursor.src.Content
	cursor.currentNode.src = cursor.src
	cursor.currentNode.Nature = nature
	return nil
}
