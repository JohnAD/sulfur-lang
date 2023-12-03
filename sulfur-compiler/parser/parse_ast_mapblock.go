package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstMapBlock(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//case lexer.TT_OPEN_SYMBOL:
	//	return openSymbolHandling(cursor, token)
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstMapBlock(cursor, token)
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//	return binderHandling(cursor, token)
	case lexer.TT_IDENT:
		return parseAstMapBlockItemStart(cursor, token)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil // just ignore when in a ROLNE (but not inside a ROLNE ITEM)
	case lexer.TT_INDENT_LINE:
		return nil
	}
	return fmt.Errorf("unhandled AST_MAPBLOCK parse of %v", token)
}

func parseAstMapBlockStart(cursor *parseCursor, token lexer.Token) error {
	createAndBecomeChild(cursor, AST_MAPBLOCK, ASTN_GROUPING, token.Content)
	return nil
}

//func parseAstMapBlockAddItem(cursor *parseCursor, token lexer.Token) error {
//	//addChild(cursor, AST_MAPBLOCK_ITEM, ASTN_KEYWORD, token.Content)
//	//return nil
//}

func finishAstMapBlock(cursor *parseCursor, token lexer.Token) error {
	if token.Content == "}}" {
		return finishAstNode(cursor)
	}
	return fmt.Errorf("[PARSE_MAPBLOCK_FAMB] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}
