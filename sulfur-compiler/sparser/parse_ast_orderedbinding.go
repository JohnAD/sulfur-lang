package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

func parseAstOrderedBinding(cursor *parseCursor, token lexer.Token) error {
	debug(AST_ORDERED_BINDING, "MAIN", cursor)
	// this state should NEVER be called. If it is called, something has gone very wrong.
	return fmt.Errorf("unhandled AST_ORDERED_BINDING parse of %v", token)
}

func binderHandlingForLastChild(cursor *parseCursor, token lexer.Token) error {
	debug(AST_ORDERED_BINDING, "BHFLC", cursor)
	err := moveToLastChild(cursor)
	if err != nil {
		return err
	}
	return binderHandlingInPlace(cursor, token)
}

func binderHandlingInPlace(cursor *parseCursor, token lexer.Token) error {
	debug(AST_ORDERED_BINDING, "BHIP", cursor)
	s := cursor.currentNode
	addChild(cursor, AST_ORDERED_BINDING_CHILD, s.Nature, s.Name, false)
	// "keep" the Kind as AST_ORDERED_BINDING (the parent) is never actually invoked
	cursor.currentNode.Nature = ASTN_META_BINDING
	cursor.currentNode.Name = token.Content
	addChild(cursor, AST_ORDERED_BINDING_CHILD, ASTN_NULL, "", false)
	return gotoChild(cursor, 1)
}
