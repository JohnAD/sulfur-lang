package sparser

import (
	"fmt"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
)

// A STATEMENT is a node with no Name and a default nature of null by default.
// All the content of a statement is in the children.

func parseAstStatement(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		return interpretInlineTokenDuringStatement(cursor)
	case lexer.TT_OPEN_SYMBOL:
		return openSymbolHandlingForNewChild(cursor)
	//case lexer.TT_CLOSE_SYMBOL:
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
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE:
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
	name := cursor.src.Content
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
	addChild(cursor, AST_STATEMENT_ITEM, nature, name)
	return nil
}

func finishStatement(cursor *parseCursor) error {
	return finishAstNode(cursor)
}
