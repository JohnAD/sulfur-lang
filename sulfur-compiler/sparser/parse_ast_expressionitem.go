package sparser

import "sulfur-compiler/lexer"

// An EXPRESSION-ITEM is a single node under an EXPRESSION that contains the "result" of the expression

func parseAstExpressionItem(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	if cursor.currentNode.Nature == ASTN_NULL {
		return parseAstExpressionItemApplyInitialToken(cursor)
	} else {
		switch cursor.src.TokenType {
		case lexer.TT_STANDING_SYMBOL:
			return startAstInfixOperandInPlace(cursor)
		}
	}
	return parseError(cursor, "MAIN", "unhandled parse")
}

func parseAstExpressionItemApplyInitialToken(cursor *parseCursor) error {
	debug(cursor, "PAEIAIT")
	switch cursor.src.TokenType {
	case lexer.TT_IDENT:
		applyNameNatureToSelf(cursor, ASTN_IDENTIFIER)
		return nil
	}
	return parseError(cursor, "PAEIAIT", "unhandled parse")
}
