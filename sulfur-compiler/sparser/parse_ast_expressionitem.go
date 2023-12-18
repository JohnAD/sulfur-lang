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
			return parseAstInfixStartInPlace(cursor)
		case lexer.TT_CLOSE_SYMBOL:
			return finishAstExpressionItem(cursor)
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
	case lexer.TT_NUMSTR_LIT:
		applyNameNatureToSelf(cursor, ASTN_NUMSTR)
		return nil
	case lexer.TT_STR_LIT:
		applyNameNatureToSelf(cursor, ASTN_STRLIT)
		return nil
	}
	return parseError(cursor, "PAEIAIT", "unhandled parse")
}

func finishAstExpressionItem(cursor *parseCursor) error {
	err := finishAstNode(cursor)
	if err != nil {
		return err
	}
	return finishAstExpression(cursor)
}
