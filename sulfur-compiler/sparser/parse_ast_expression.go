package sparser

import (
	"fmt"
	"sulfur-compiler/helpers"
)

// An EXPRESSION is a node with no Name and a series of EXPRESSION_ITEM children in the order of expression.
// The expression is NOT evaluated in any way as SPARSER does not handle order-of-operation
// https://en.wikipedia.org/wiki/Order_of_operations
// Further grouping by parens and the like can create expressions under expressions.
//
// This parental node contains the "opening" token of the expression (IF there is one.) For example, "(".
//
// example:
//
//         ( 1 + 2 * 3)
//
//     becomes:
//
//         EXPRESSION("(") -> [
//           EXPRESSION-ITEM(1)
//           EXPRESSION-ITEM(+)
//           EXPRESSION-ITEM(2)
//           EXPRESSION-ITEM(*)
//           EXPRESSION-ITEM(3)
//         ]
//
// another:
//
//         ( ( 1 + 2 ) * 3)
//
//     becomes:
//
//         EXPRESSION("(") -> [
//           EXPRESSION("(") -> [
//             EXPRESSION-ITEM(1)
//             EXPRESSION-ITEM(+)
//             EXPRESSION-ITEM(2)
//           ]
//           EXPRESSION-ITEM(*)
//           EXPRESSION-ITEM(3)
//         ]
//
// example:
//
//         if ( 1 == 3 ) then
//
//     becomes
//
//         STATEMENT() -> [
//            S-ITEM(if)
//            S-ITEM() -> [
//              EXPRESSION("(") -> [
//                EXPRESSION-ITEM(1)
//                EXPRESSION-ITEM(==)
//                EXPRESSION-ITEM(3)
//              ]
//            ]
//            S-ITEM(then)
//
//
// FOR NOW, expressions require the parens (). But later, the parse will be expanded to not require it:
//
// example:
//
//         if 1 == 3 then
//
//     becomes
//
//         STATEMENT() -> [
//            S-ITEM(if)
//            S-ITEM() -> [
//              EXPRESSION() -> [
//                EXPRESSION-ITEM(1)
//                EXPRESSION-ITEM(==)
//                EXPRESSION-ITEM(3)
//              ]
//            ]
//            S-ITEM(then)
//
// this works by having the parser detect the STANDING OP (==) and trigger the expression with the previous (1).
// the expression would continue to build until the <non-op> <op> <non-op> <op> ... pattern is broken.
//

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

func finishAstExpression(cursor *parseCursor) error {
	debug(cursor, "FAI")
	starting := cursor.currentNode.Src.Content
	ending := cursor.src.Content
	if helpers.OperatorsMatch(starting, ending) {
		err := finishAstNode(cursor)
		debug(cursor, "FAI")
		return err
	}
	return parseError(cursor, "FAI", fmt.Sprintf("unable to match '%s' with '%s'", starting, ending))
}
