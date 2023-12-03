package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

// every AST_MAPBLOCK_ITEM should have EXACTLY TWO ITEMS: identifier target

func parseAstMapBlockItem(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	case lexer.TT_OPEN_SYMBOL:
		return openSymbolHandling(cursor, token)
	//case lexer.TT_CLOSE_SYMBOL:
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//case lexer.TT_IDENT:
	//	return parseAstMapBlockItemValue(cursor, token)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil
	case lexer.TT_INDENT_LINE:
		return finishAstMapBlockItem(cursor)
	}
	return fmt.Errorf("unhandled AST_MAPBLOCK_ITEM parse of %v", token)
}

func parseAstMapBlockItemStart(cursor *parseCursor, token lexer.Token) error {
	createAndBecomeChild(cursor, AST_MAPBLOCK_ITEM, ASTN_GROUPING, token.Content)
	return nil
}

func finishAstMapBlockItem(cursor *parseCursor) error {
	return finishAstNode(cursor)
}

//func parseAstMapBlockItemValue(cursor *parseCursor, token lexer.Token) error {
//	addChild(cursor, AST_IDENTIFIER, ASTN_IDENTIFIER, token.Content)
//	return nil
//}
