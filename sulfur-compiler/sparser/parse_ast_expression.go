package sparser

// An EXPRESSION is a node with no Name and a single EXPRESSION_ITEM child that contains
// the result of the expression.
//
// This parental node contains the "opening" token of the expression (if there is one.) For example, "(".

func parseAstExpression(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	}
	return parseError(cursor, "MAIN", "unhandled parse")
}

func parseAstExpressionStartChild(cursor *parseCursor) error {
	debug(cursor, "START_CHILD")
	createAndBecomeChild(cursor, AST_EXPRESSION, ASTN_META_BINDING)
	createAndBecomeEmptyChild(cursor, AST_EXPRESSION_ITEM, ASTN_NULL)
	return nil
}
