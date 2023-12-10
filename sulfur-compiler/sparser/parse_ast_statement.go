package sparser

import (
	"fmt"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
)

// A STATEMENT is a node with no Name and a default nature of null by default.
// All the content of a statement is in the children.

func parseAstStatement(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		return interpretInlineTokenDuringStatement(cursor, token)
	case lexer.TT_OPEN_SYMBOL:
		return openSymbolHandlingForNewChild(cursor, token)
	//case lexer.TT_CLOSE_SYMBOL:
	case lexer.TT_OPEN_BIND_SYMBOL:
		return openSymbolHandlingForLastChild(cursor, token, true)
	case lexer.TT_BINDING_SYMBOL:
		return binderHandling(cursor, token)
	case lexer.TT_IDENT:
		return interpretInlineTokenDuringStatement(cursor, token)
	case lexer.TT_STR_LIT:
		return interpretInlineTokenDuringStatement(cursor, token)
	case lexer.TT_NUMSTR_LIT:
		return interpretInlineTokenDuringStatement(cursor, token)
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE:
	case lexer.TT_LINE:
		return finishStatement(cursor, token)
	}
	return fmt.Errorf("unhandled AST_STATEMENT parse of %v", token)
}

func parseAstStatementStartChild(cursor *parseCursor, token lexer.Token) error {
	nature := interpretInitialStatementNature(token)
	createAndBecomeChild(cursor, AST_STATEMENT, nature, "", false)
	// make one call to main proc to interpret the first child
	return parseAstStatement(cursor, token)
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

func interpretInlineTokenDuringStatement(cursor *parseCursor, token lexer.Token) error {
	nature := ASTN_IDENTIFIER
	name := token.Content
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if helpers.Contains([]string{"#!", "#@"}, token.Content) {
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
		return fmt.Errorf("[PARSE_STMT_IITDS] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
	}
	addChild(cursor, AST_STATEMENT_ITEM, nature, name, false)
	return nil
}

func finishStatement(cursor *parseCursor, token lexer.Token) error {
	return finishAstNode(cursor)
}
