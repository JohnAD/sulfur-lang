package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstRolne(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		return parseAstRolnePossibleAssignment(cursor, token)
	//case lexer.TT_OPEN_SYMBOL:
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstRolne(cursor, token)
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	case lexer.TT_IDENT:
		return parseAstRolneItemSimple(cursor, token)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	case lexer.TT_COMMENT:
		return nil
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil // just ignore when in a ROLNE (but not inside a ROLNE ITEM)
	case lexer.TT_INDENT_LINE:
		return nil // just ignore EOL when in a ROLNE (but not inside a ROLNE ITEM)
	}
	return fmt.Errorf("unhandled AST_ROLNE parse of %v", token)
}

func parseAstRolneItemSimple(cursor *parseCursor, token lexer.Token) error {
	addChild(cursor, AST_ROLNE_ITEM, ASTN_IDENTIFIER, token.Content)
	return nil
}

func parseAstRolnePossibleAssignment(cursor *parseCursor, token lexer.Token) error {
	if token.Content == "=" {
		return parseAstRolneItemStartAssignment(cursor, token)
	}
	return fmt.Errorf("[PARSE_ROLNE_PARPA] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}

func parseAstRolneStartChild(cursor *parseCursor, token lexer.Token) error {
	createAndBecomeChild(cursor, AST_ROLNE, ASTN_META_BINDING, token.Content)
	return nil
}

func finishAstRolne(cursor *parseCursor, token lexer.Token) error {
	if token.Content == "}" {
		return finishAstNode(cursor)
	}
	return fmt.Errorf("[PARSE_ROLNE_FAR] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}
