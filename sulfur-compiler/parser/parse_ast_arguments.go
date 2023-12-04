package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

// a `ident( )` is a list of arguments, generally for a call (but also used for types etc.)

func parseAstArguments(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//case lexer.TT_OPEN_SYMBOL:
	//	return openSymbolHandlingForNewChild(cursor, token)
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstArguments(cursor, token)
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//	return binderHandling(cursor, token)
	case lexer.TT_IDENT:
		return parseAstArgumentsNextItem(cursor, token, ASTN_IDENTIFIER)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	case lexer.TT_LINE: // just ignore when in ARGUMENTS
		return nil
	}
	return fmt.Errorf("unhandled AST_ARGUMENTS parse of %v", token)
}

func parseAstArgumentsStartChild(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	createAndBecomeChild(cursor, AST_ARGUMENTS, nature, token.Content)
	return nil
}

func parseAstArgumentsStart(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	selfPtr := cursor.currentNode
	selfPtr.Kind = AST_ARGUMENTS
	selfPtr.Nature = nature
	selfPtr.Name = token.Content
	return nil
}

func finishAstArguments(cursor *parseCursor, token lexer.Token) error {
	if token.Content == ")" {
		_ = finishAstNode(cursor)
		return finishAstNode(cursor)
	}
	return fmt.Errorf("[PARSE_ARGS_FAA] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}

func parseAstArgumentsNextItem(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	addChild(cursor, AST_ARGUMENT_ITEM, nature, token.Content)
	return nil
}
