package sparser

import "sulfur-compiler/lexer"

// An EXPRESSION-ITEM is a single node under an EXPRESSION
// see EXPRESSION for more detail

func parseAstExpressionItem(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	if cursor.currentNode.Nature == ASTN_NULL {
		return parseAstExpressionItemApplyInitialToken(cursor)
	} else {
		switch cursor.src.TokenType {
		case lexer.TT_IDENT:
			return parseAstExpressionItemBecomeNextChild(cursor, ASTN_IDENTIFIER)
		case lexer.TT_NUMSTR_LIT:
			return parseAstExpressionItemBecomeNextChild(cursor, ASTN_NUMSTR)
		case lexer.TT_STR_LIT:
			return parseAstExpressionItemBecomeNextChild(cursor, ASTN_STRLIT)
		case lexer.TT_STANDING_SYMBOL:
			return parseAstExpressionItemBecomeNextChild(cursor, ASTN_INFIX_OPERATOR)
		case lexer.TT_CLOSE_SYMBOL:
			return parseAstExpressionItemFinish(cursor)
		default:
			return parseError(cursor, "MAIN", "unhandled token type")
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

func parseAstExpressionItemBecomeNextChild(cursor *parseCursor, nature AstNodeNature) error {
	debug(cursor, "PAEIBNC")
	err := finishAstNode(cursor)
	if err != nil {
		return err
	}
	createAndBecomeChild(cursor, AST_EXPRESSION_ITEM, nature)
	return nil
}

func parseAstExpressionItemFinish(cursor *parseCursor) error {
	debug(cursor, "PAEIF")
	err := finishAstNode(cursor)
	if err != nil {
		return err
	}
	return finishAstExpression(cursor)
}
