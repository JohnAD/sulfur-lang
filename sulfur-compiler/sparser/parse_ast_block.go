package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

// a [[ ]] block holds a list of statements. There is no return value from the block.

func parseAstBlock(cursor *parseCursor) error {
	switch cursor.src.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//case lexer.TT_OPEN_SYMBOL:
	//	return openBindSymbolHandlingForNewChild(cursor)
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstBlock(cursor)
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//	return binderHandlingForLastChild(cursor)
	case lexer.TT_IDENT:
		return parseAstStatementStartChild(cursor)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	case lexer.TT_LINE: // just ignore when in a ROLNE (but not inside a ROLNE ITEM)
		return nil
	}
	return fmt.Errorf("unhandled AST_BLOCK parse of %v", cursor.src)
}

func parseAstBlockStartChild(cursor *parseCursor, nature AstNodeNature) error {
	createAndBecomeChild(cursor, AST_BLOCK, nature)
	return nil
}

func parseAstBlockStart(cursor *parseCursor, nature AstNodeNature) error {
	selfPtr := cursor.currentNode
	selfPtr.Kind = AST_BLOCK
	selfPtr.Nature = nature
	selfPtr.Src = cursor.src
	return nil
}

func finishAstBlock(cursor *parseCursor) error {
	if cursor.src.Content == "]]" {
		return finishAstNode(cursor)
	}
	return fmt.Errorf("[PARSE_BLOCK_FAB] unable to determine what '%s' is on line %d column %d", cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset)
}
