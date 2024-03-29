package sparser

import (
	"fmt"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
)

// A STATEMENT is a node with no Name and a default nature of null by default.
// All the content of a statement is in the children who are ALL S-ITEM nodes, some with children of their own.
//
// example:
//
//         if ( 1 == 3 ) [[
//           clib\print( "equal" )
//         ]] else [[
//           clib\print( "not equal" )
//         ]]
//
//     becomes a statement with 5 items (only ast and token's content shown)
//
//         STATEMENT() -> [
//           S-ITEM(if)                                // #1
//           S-ITEM() -> [                             // #2
//             EXPRESSION("(") -> [
//               EXPRESSION-ITEM(1)
//               EXPRESSION-ITEM(==)
//               EXPRESSION-ITEM(3)
//             ]
//           ]
//           S-ITEM() -> [                             // #3
//             BLOCK("[[") -> [
//               STATEMENT() -> [
//                 S-ITEM() -> [
//                   BINDING() -> [
//                     B-ITEM(clib)
//                     B-ITEM("\")
//                     B-ITEM(print) -> [
//                       ROLNE("(") -> [
//                         R-ITEM() -> [
//                           R-ITEM-NAME()
//                           R-ITEM-TYPE()
//                           R-ITEM-VALUE(equal)
//                         ]
//                       ]
//                     ]
//                   ]
//                 ]
//               ]
//             ]
//           ]
//           S-ITEM(else)                              // #4
//           S-ITEM() -> [                             // #5
//             BLOCK("[[") -> [
//               STATEMENT() -> [
//                 S-ITEM() -> [
//                   BINDING() -> [
//                     B-ITEM(clib)
//                     B-ITEM("\")
//                     B-ITEM(print) -> [
//                       ROLNE("(") -> [
//                         R-ITEM() -> [
//                           R-ITEM-NAME()
//                           R-ITEM-TYPE()
//                           R-ITEM-VALUE("not equal")
//                         ]
//                       ]
//                     ]
//                   ]
//                 ]
//               ]
//             ]
//           ]
//         ]
//
//
// TODO: add the above as a unit test

func parseAstStatement(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		return interpretInlineTokenDuringStatement(cursor)
	case lexer.TT_OPEN_SYMBOL:
		return openSymbolHandlingForNewChild(cursor)
	case lexer.TT_OPEN_BIND_SYMBOL:
		return openSymbolHandlingForLastChild(cursor)
	case lexer.TT_BINDING_SYMBOL:
		return binderHandlingForLastChild(cursor)
	case lexer.TT_IDENT:
		return interpretInlineTokenDuringStatement(cursor)
	case lexer.TT_STR_LIT:
		return interpretInlineTokenDuringStatement(cursor)
	case lexer.TT_NUMSTR_LIT:
		return interpretInlineTokenDuringStatement(cursor)
	case lexer.TT_LINE:
		return finishStatement(cursor)
	}
	return fmt.Errorf("unhandled AST_STATEMENT parse of %v", cursor.src)
}

func parseAstStatementStartChild(cursor *parseCursor) error {
	nature := interpretInitialStatementNature(cursor.src)
	createAndBecomeEmptyChild(cursor, AST_STATEMENT, nature)
	// make one call to main proc to interpret the first child
	return parseAstStatement(cursor)
}

func interpretInitialStatementNature(token lexer.Token) AstNodeNature {
	if token.SourceLine == 1 {
		if token.Content == "#!" {
			return ASTN_STATEMENT_ROOT_DECLARATION
		}
	}
	if token.SourceLine == 2 {
		if token.Content == "#@" {
			return ASTN_STATEMENT_ROOT_FRAMEWORK
		}
	}
	return ASTN_NULL
}

func interpretInlineTokenDuringStatement(cursor *parseCursor) error {
	debug(cursor, "IITDS")
	nature := ASTN_IDENTIFIER
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if helpers.Contains([]string{"#!", "#@"}, cursor.src.Content) {
			nature = ASTN_KEYWORD
		} else {
			nature = ASTN_INFIX_OPERATOR
		}
	case lexer.TT_IDENT:
		nature = ASTN_IDENTIFIER
	case lexer.TT_STR_LIT:
		nature = ASTN_STRLIT
	case lexer.TT_NUMSTR_LIT:
		nature = ASTN_NUMSTR
	default:
		return fmt.Errorf("[PARSE_STMT_IITDS] unable to determine what '%s' is on line %d column %d", cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset)
	}
	addChild(cursor, AST_STATEMENT_ITEM, nature)
	return nil
}

func finishStatement(cursor *parseCursor) error {
	return finishAstNode(cursor)
}
