package sparser

import "sulfur-compiler/lexer"

// An INFIX_RIGHT is the preceding node before an INFIX operator.
//
// For example, 1 + 2 becomes INFIX(+) with children INFIX_LEFT(1) and INFIX_RIGHT(2)
//

func parseAstInfixRight(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	case lexer.TT_NUMSTR_LIT:
		return parseAstInfixRightFinish(cursor, ASTN_NUMSTR)
	}
	return parseError(cursor, "MAIN", "unhandled parse")
}

func parseAstInfixRightFinish(cursor *parseCursor, nature AstNodeNature) error {
	applyNameNatureToSelf(cursor, nature)
	return finishAstNode(cursor) // finish AST_INFIX_RIGHT
}
