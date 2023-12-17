package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstMapBlock(cursor *parseCursor) error {
	switch cursor.src.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//case lexer.TT_OPEN_SYMBOL:
	//	return openSymbolHandlingForNewChild(cursor)
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstMapBlock(cursor)
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//	return binderHandlingForLastChild(cursor)
	case lexer.TT_IDENT:
		return parseAstMapBlockItemStart(cursor)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil // just ignore when in a ROLNE (but not inside a ROLNE ITEM)
	case lexer.TT_LINE:
		return nil
	}
	return fmt.Errorf("unhandled AST_MAPBLOCK parse of %v", cursor.src)
}

func parseAstMapBlockStart(cursor *parseCursor) error {
	createAndBecomeChild(cursor, AST_MAPBLOCK, ASTN_GROUPING)
	return nil
}

//func parseAstMapBlockAddItem(cursor *parseCursor) error {
//	//addChild(cursor, AST_MAPBLOCK_ITEM, ASTN_KEYWORD, token.Content)
//	//return nil
//}

func finishAstMapBlock(cursor *parseCursor) error {
	if cursor.src.Content == "}}" {
		return finishAstNode(cursor)
	}
	return fmt.Errorf("[PARSE_MAPBLOCK_FAMB] unable to determine what '%s' is on line %d column %d", cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset)
}
