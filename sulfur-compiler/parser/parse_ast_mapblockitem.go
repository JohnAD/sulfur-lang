package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

// TODO: verify below
// every AST_MAPBLOCK_ITEM should have EXACTLY TWO ITEMS: identifier target
// in practice, this means each each MT-ITEM has a name (the identifier) and one child (the target)

func parseAstMapBlockItem(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	case lexer.TT_OPEN_SYMBOL:
		err := gotoChild(cursor, 0)
		if err != nil {
			return fmt.Errorf("[PARSE_MAPBLOCKITEM_OPENSYMBOL] %v", err)
		}
		return openSymbolHandlingInPlace(cursor, token)
	//case lexer.TT_CLOSE_SYMBOL:
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//case lexer.TT_IDENT:
	//	return parseAstMapBlockItemValue(cursor, token)
	case lexer.TT_STR_LIT:
		return setChild(cursor, 0, AST_LITERAL, ASTN_STR, token.Content)
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil
	case lexer.TT_LINE:
		return finishAstMapBlockItem(cursor, token)
	}
	return fmt.Errorf("unhandled AST_MAPBLOCK_ITEM parse of %v", token)
}

func parseAstMapBlockItemStart(cursor *parseCursor, token lexer.Token) error {
	createAndBecomeChild(cursor, AST_MAPBLOCK_ITEM, ASTN_IDENTIFIER, token.Content, false)
	addChild(cursor, AST_TARGET, ASTN_NULL, "", false)
	return nil
}

func finishAstMapBlockItem(cursor *parseCursor, token lexer.Token) error {
	if cursor.currentNode.Nature == ASTN_NULL {
		return fmt.Errorf("[PARSE_MAPBLOCKITEM_FAMBI] identifier '%s' found, but no matching target; line %d column %d", cursor.currentNode.Name, token.SourceLine, token.SourceOffset)
	}
	return finishAstNode(cursor)
}
