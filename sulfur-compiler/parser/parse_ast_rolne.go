package parser

import (
	"fmt"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
)

func parseAstRolne(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	//case lexer.TT_STANDING_SYMBOL:
	//	return parseAstRolnePossibleAssignment(cursor, token)
	//case lexer.TT_OPEN_SYMBOL:
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstRolne(cursor, token)
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//	return binderHandling(cursor, token)
	case lexer.TT_IDENT:
		return parseAstRolneItemStart(cursor, token, ASTN_IDENTIFIER, false)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//	return nil
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil // just ignore when in a ROLNE (but not inside a ROLNE ITEM)
	case lexer.TT_LINE:
		return nil // just ignore EOL when in a ROLNE (but not inside a ROLNE ITEM)
	}
	return fmt.Errorf("unhandled AST_ROLNE parse of %v", token)
}

func parseAstRolneStart(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	selfPtr := cursor.currentNode
	selfPtr.Kind = AST_ROLNE
	selfPtr.Nature = nature
	selfPtr.Name = token.Content
	return nil
}

func parseAstRolneStartChild(cursor *parseCursor, token lexer.Token, nature AstNodeNature, bound bool) error {
	createAndBecomeChild(cursor, AST_ROLNE, nature, token.Content, bound)
	return nil
}

func finishAstRolne(cursor *parseCursor, token lexer.Token) error {
	starting := cursor.currentNode.Name
	ending := token.Content
	if helpers.OperatorsMatch(starting, ending) {
		err := finishAstNode(cursor)
		return err
	}
	return fmt.Errorf("[PARSE_ROLNE_FAR] unable to match '%' with '%s' on line %d column %d", starting, ending, token.SourceLine, token.SourceOffset)
}
