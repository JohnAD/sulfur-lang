package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolneItemName(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if token.Content == "=" {
			return parseAstRolneItemValueStartViaEqualSign(cursor)
		} else if token.Content == "," {
			return childParseAstRolneItemFinish(cursor, token)
		}
	case lexer.TT_OPEN_BIND_SYMBOL:
		return openSymbolHandlingForNewChild(cursor, token)
	case lexer.TT_CLOSE_SYMBOL:
		return childParseAstRolneItemFinish(cursor, token) // a closing "}" etc found
	case lexer.TT_IDENT:
		return parseAstRolneItemNameAssignName(cursor, token, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemNameAssignName(cursor, token, ASTN_STRLIT)
	case lexer.TT_LINE:
		return childParseAstRolneItemFinish(cursor, token)
	}
	return fmt.Errorf("unhandled AST_ROLNE_NAME parse of %v", token)
}

func parseAstRolneItemNameStart(cursor *parseCursor, token lexer.Token) error {
	// when this is called, we should ALREADY be pointing to the name node
	debug("PARINS", cursor)
	return parseAstRolneItemName(cursor, token)
}

func parseAstRolneItemNameAssignName(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	if cursor.currentNode.Nature != ASTN_NULL {
		// if a "second name" is found, it means the current item is a value and a new rolne item is starting
		return childParseAstRolneItemFinish(cursor, token)
	}
	cursor.currentNode.Name = token.Content
	cursor.currentNode.Nature = nature
	return nil
}
