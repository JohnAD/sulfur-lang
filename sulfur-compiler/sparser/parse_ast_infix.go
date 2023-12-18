package sparser

// An INFIX is an operator node with INFIX_LEFT_ITEM and INFIX_RIGHT_ITEM children.
//
// For example, 1 + 2 becomes INFIX(+) with children INFIX_LEFT(1) and INFIX_RIGHT(2)
//

const AST_INFIX_LEFT_INDEX = 0
const AST_INFIX_RIGHT_INDEX = 1

func parseAstInfix(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	}
	return parseError(cursor, "MAIN", "unhandled parse")
}

func parseAstInfixStartInPlace(cursor *parseCursor) error {
	// seq:
	//  1. the "current" node becomes the left child node
	//  2. the token becomes current
	//  3. the current then points to the right
	debug(cursor, "START_IN_PLACE")
	addChildWithCopyOfCurrent(cursor, AST_INFIX_LEFT)
	addNullChild(cursor, AST_INFIX_RIGHT)
	applyNameNatureToSelf(cursor, ASTN_INFIX_OPERATOR)
	return gotoChild(cursor, AST_INFIX_RIGHT_INDEX)
}
