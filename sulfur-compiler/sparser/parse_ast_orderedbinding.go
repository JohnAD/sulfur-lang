package sparser

import (
	"fmt"
)

func parseAstOrderedBinding(cursor *parseCursor) error {
	debug(AST_ORDERED_BINDING, "MAIN", cursor)
	// this state should NEVER be called. If it is called, something has gone very wrong.
	return fmt.Errorf("unhandled AST_ORDERED_BINDING parse of %v", cursor.src)
}

func binderHandlingForLastChild(cursor *parseCursor) error {
	debug(AST_ORDERED_BINDING, "BHFLC", cursor)
	err := moveToLastChild(cursor)
	if err != nil {
		return err
	}
	return binderHandlingInPlace(cursor)
}

func binderHandlingInPlace(cursor *parseCursor) error {
	debug(AST_ORDERED_BINDING, "BHIP", cursor)
	s := cursor.currentNode
	addChild(cursor, AST_ORDERED_BINDING_CHILD, s.Nature, s.Name)
	// "keep" the Kind as AST_ORDERED_BINDING (the parent) is never actually invoked
	cursor.currentNode.Nature = ASTN_META_BINDING
	cursor.currentNode.Name = cursor.src.Content
	cursor.currentNode.src = cursor.src
	addChild(cursor, AST_ORDERED_BINDING_CHILD, ASTN_NULL, "")
	return gotoChild(cursor, 1)
}
