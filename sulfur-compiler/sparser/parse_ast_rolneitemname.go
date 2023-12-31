package sparser

import (
	"sulfur-compiler/lexer"
)

func parseAstRolneItemName(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if cursor.src.Content == "=" {
			return parseAstRolneItemValueStartViaEqualSign(cursor)
		} else if cursor.src.Content == "," {
			return childParseAstRolneItemFinish(cursor)
		}
	case lexer.TT_OPEN_BIND_SYMBOL:
		return openBindSymbolHandlingForNewChild(cursor)
	case lexer.TT_BINDING_SYMBOL:
		// any form of binding says "this isn't a name" since names CANNOT be non-simple per language rules
		if cursor.src.Content == "::" {
			return parseAstRolneItemTypeStart(cursor)
		} else {
			return parseAstRolneItemValueStartViaBinding(cursor)
		}
	case lexer.TT_CLOSE_SYMBOL:
		return childParseAstRolneItemFinish(cursor) // a closing "}" etc found
	case lexer.TT_IDENT:
		return parseAstRolneItemNameAssignName(cursor, cursor.src, ASTN_IDENTIFIER)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemNameAssignName(cursor, cursor.src, ASTN_STRLIT)
	case lexer.TT_NUMSTR_LIT:
		return parseAstRolneItemNameAssignName(cursor, cursor.src, ASTN_NUMSTR)
	case lexer.TT_LINE:
		return childParseAstRolneItemFinish(cursor)
	}
	return parseError(cursor, "MAIN", "unhandled parse")
}

func parseAstRolneItemNameStart(cursor *parseCursor) error {
	// when this is called, we should ALREADY be pointing to the name node
	debug(cursor, "PARINS")
	return parseAstRolneItemName(cursor)
}

func parseAstRolneItemNameAssignName(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	if cursor.currentNode.Nature != ASTN_NULL {
		// if a "second name" is found, it means the current item is a value and a new rolne item is starting
		return childParseAstRolneItemFinish(cursor)
	}
	cursor.currentNode.Src = token
	cursor.currentNode.Nature = nature
	return nil
}
