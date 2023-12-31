package sparser

import (
	"fmt"
)

// A BINDING is a node with no Name or Token and an ordered series of B-ITEM children in the order of the binding.
//
// example:
//
//         a\b\c.d
//
//     becomes:
//
//         BINDING() -> [
//           B-ITEM(a)
//           B-ITEM(\)
//           B-ITEM(b)
//           B-ITEM(\)
//           B-ITEM(c)
//           B-ITEM(.)
//           B-ITEM(d)
//         ]
//
// At this stage "resolving" these bindings is not possible yet as we don't see the cross-refs of libraries and such.
// TODO: change expression to actually do what the above says

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
