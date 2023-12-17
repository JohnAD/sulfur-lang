package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolneItemType(cursor *parseCursor, token lexer.Token) error {
	debug(AST_ROLNE_ITEM_TYPE, "MAIN", cursor)
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if token.Content == "=" {
			return parseAstRolneItemValueStartViaEqualSign(cursor)
		} else if token.Content == "," {
			return childParseAstRolneItemFinish(cursor, token)
		}
	case lexer.TT_CLOSE_SYMBOL:
		return childParseAstRolneItemFinish(cursor, token) // a closing "}" etc. found
	case lexer.TT_IDENT:
		return parseAstRolneItemTypeAssignType(cursor, token, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemTypeAssignType(cursor, token, ASTN_STRLIT)
	case lexer.TT_LINE:
		return childParseAstRolneItemFinish(cursor, token)
	}
	return fmt.Errorf("unhandled AST_ROLNE_ITEM_VALUE parse of %v", token)
}

func parseAstRolneItemTypeStart(cursor *parseCursor) error {
	// when this is called, we should be currently pointing to the AST_ROLNE_ITEM_NAME
	debug(AST_ROLNE_ITEM_VALUE, "PARITS", cursor)
	_ = finishAstNode(cursor)                    // finish and point to AST_ROLNE_ITEM
	return gotoChild(cursor, ROLEITEM_TYPECHILD) // now point to AST_ROLNE_ITEM_TYPE
}

func parseAstRolneItemTypeAssignType(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	debug(AST_ROLNE_ITEM_TYPE, "PARITAV-start", cursor)
	if cursor.currentNode.Nature != ASTN_NULL {
		return childParseAstRolneItemFinish(cursor, token)
	}
	cursor.currentNode.Name = token.Content
	cursor.currentNode.src = token
	cursor.currentNode.Nature = nature
	return nil
}
