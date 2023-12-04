package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstStatement(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		return interpretInlineTokenDuringStatement(cursor, token)
	case lexer.TT_OPEN_SYMBOL:
		return openSymbolHandlingForNewChild(cursor, token)
	//case lexer.TT_CLOSE_SYMBOL:
	case lexer.TT_OPEN_BIND_SYMBOL:
		return openSymbolHandlingForLastChild(cursor, token)
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
	err, nature, name := interpretTokenAtStartOfStatement(token)
	createAndBecomeChild(cursor, AST_STATEMENT, nature, name)
	return err
}

func interpretTokenAtStartOfStatement(token lexer.Token) (error, AstNodeNature, string) {
	if token.SourceLine == 1 {
		if token.Content == "#!" {
			return nil, ASTN_STATEMENT_ROOT_DECLARATION, token.Content
		} else {
			return fmt.Errorf("root declaration ('#!') missing on line %d", token.SourceLine), 0, ""
		}
	}
	if token.SourceLine == 2 {
		if token.Content == "#@" {
			return nil, ASTN_STATEMENT_ROOT_FRAMEWORK, token.Content
		}
	}
	switch token.TokenType {
	case lexer.TT_IDENT:
		return nil, ASTN_IDENTIFIER, token.Content
	}
	return fmt.Errorf("[PARSE_STMT_ITASOS] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset), 0, ""
}

func interpretInlineTokenDuringStatement(cursor *parseCursor, token lexer.Token) error {
	nature := ASTN_IDENTIFIER
	name := token.Content
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		nature = ASTN_INFIX_OPERATOR
	case lexer.TT_IDENT:
		nature = ASTN_IDENTIFIER
	case lexer.TT_STR_LIT:
		nature = ASTN_STR
	case lexer.TT_NUMSTR_LIT:
		nature = ASTN_NUMSTR
	default:
		return fmt.Errorf("[PARSE_STMT_IITDS] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
	}
	addChild(cursor, AST_STATEMENT_ITEM, nature, name)
	return nil
}

func finishStatement(cursor *parseCursor, token lexer.Token) error {
	return finishAstNode(cursor)
}
