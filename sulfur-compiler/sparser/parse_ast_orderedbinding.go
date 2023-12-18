package sparser

import (
	"fmt"
)

func parseAstOrderedBinding(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	// this state should NEVER be called. If it is called, something has gone very wrong.
	return fmt.Errorf("unhandled AST_ORDERED_BINDING parse of %v", cursor.src)
}

func binderHandlingForLastChild(cursor *parseCursor) error {
	debug(cursor, "BHFLC")
	err := moveToLastChild(cursor)
	if err != nil {
		return err
	}
	return binderHandlingInPlace(cursor)
}

func binderHandlingInPlace(cursor *parseCursor) error {
	debug(cursor, "BHIP")
	addChildWithCopyOfCurrent(cursor, AST_ORDERED_BINDING_CHILD) // child #1
	addNullChild(cursor, AST_ORDERED_BINDING_CHILD)              // child #2
	applyNameNatureToSelf(cursor, ASTN_META_BINDING)
	return gotoChild(cursor, 1) // goto child #2
}
