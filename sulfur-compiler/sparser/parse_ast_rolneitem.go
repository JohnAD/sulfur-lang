package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

// type RoleItemState int // for answer the question: where are we?
// const (
//
//	RIS_NAME RoleItemState = iota
//	RIS_TYPE
//	RIS_VALUE
//
// )
const ROLEITEM_NAMECHILD = 0
const ROLEITEM_TYPECHILD = 1
const ROLEITEM_VALUECHILD = 2

// every AST_ROLNE_ITEM should have EXACTLY THREE ITEMS: name type value
//   the ROLNE_ITEM itself holds the name as it's name; a name is either an identifier or a string
//   on creation, the NAME, TYPE and VALUE children are created pre-emptively and the current pointer moves to the name.

func parseAstRolneItem(cursor *parseCursor, token lexer.Token) error {
	// if we somehow manage to be in this state, then we are "between" rolne items
	//   on entry we add children and point to the AST_ROLNE_NAME child
	//   on exit we should pop back up to parent AST_ROLNE
	return fmt.Errorf("unhandled AST_ROLNE_ITEM parse of %v", token)
}

func parseAstRolneItemStart(cursor *parseCursor, token lexer.Token) error {
	debug("PARIS-start", cursor)
	createAndBecomeChild(cursor, AST_ROLNE_ITEM, ASTN_NOTHING, "", false) // create/become R-ITEM
	addChild(cursor, AST_ROLNE_NAME, ASTN_NULL, "", false)                // add yet-unknown name
	addChild(cursor, AST_ROLNE_TYPE, ASTN_NULL, "", false)                // add yet-unknown type
	addChild(cursor, AST_ROLNE_VALUE, ASTN_NULL, "", false)               // add yet-unknown value
	_ = gotoChild(cursor, ROLEITEM_NAMECHILD)                             // become name
	debug("PARIS-end", cursor)
	return parseAstRolneItemNameStart(cursor, token)
}

func childCloseRolneItemWithJustValue(cursor *parseCursor) {
	// a unnamed value is found, so the rolne item is now done.
	// ^ I cannot believe you made such a simple spelling mistake. You call yourself a software developer? Shame.
	// this should only be called from AST_ROLNE_NAME
	debug("CCRIWJV-start", cursor)
	formerName := cursor.currentNode.Name
	formerNature := cursor.currentNode.Nature
	formerChildren := cursor.currentNode.Children
	cursor.currentNode.Nature = ASTN_NOTHING
	cursor.currentNode.Name = ""
	cursor.currentNode.Children = []*AstNode{}
	_ = finishAstNode(cursor) // close child and point to here
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Name = formerName
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Nature = formerNature
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Children = formerChildren
	debug("CCRIWJV-end", cursor)
}

func childParseAstRolneItemFinish(cursor *parseCursor, token lexer.Token) error {
	// only be called from AST_ROLNE_NAME, AST_ROLNE_TYPE, or AST_ROLNE_VALUE
	debug("CPARIF-start", cursor)

	if cursor.currentNode.Kind == AST_ROLNE_NAME {
		childCloseRolneItemWithJustValue(cursor) // re-arrange and point to here
	}
	debug("CPARIF-end1", cursor)
	_ = finishAstNode(cursor) // point to parent AST_ROLNE
	debug("CPARIF-end2", cursor)
	return parseAstRolne(cursor, token) // have the parent handle the new token
}
