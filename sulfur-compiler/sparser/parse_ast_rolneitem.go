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

func parseAstRolneItem(cursor *parseCursor) error {
	// if we somehow manage to be in this state, then we are "between" rolne items
	//   on entry we add children and point to the AST_ROLNE_ITEM_NAME child
	//   on exit we should pop back up to parent AST_ROLNE
	return fmt.Errorf("unhandled AST_ROLNE_ITEM parse of %v", cursor.src)
}

func parseAstRolneItemStart(cursor *parseCursor) error {
	debug(AST_ROLNE_ITEM, "PARIS-start", cursor)
	createAndBecomeEmptyChild(cursor, AST_ROLNE_ITEM, ASTN_NOTHING) // create/become R-ITEM
	addChild(cursor, AST_ROLNE_ITEM_NAME, ASTN_NULL, "")            // add yet-unknown name
	addChild(cursor, AST_ROLNE_ITEM_TYPE, ASTN_NULL, "")            // add yet-unknown type
	addChild(cursor, AST_ROLNE_ITEM_VALUE, ASTN_NULL, "")           // add yet-unknown value
	_ = gotoChild(cursor, ROLEITEM_NAMECHILD)                       // become name
	debug(AST_ROLNE_ITEM, "PARIS-end", cursor)
	return parseAstRolneItemNameStart(cursor)
}

func childCloseRolneItemWithJustValue(cursor *parseCursor) {
	// bling unnamed value is found, so the rolne item is now done.
	// ^ I cannot believe you made such a simple spelling mistake. You call yourself a software developer? Shame.
	// this should only be called from AST_ROLNE_ITEM_NAME
	debug(AST_ROLNE_ITEM, "CCRIWJV-start", cursor)
	formerName := cursor.currentNode.Name
	formerSrc := cursor.currentNode.src
	formerNature := cursor.currentNode.Nature
	formerChildren := cursor.currentNode.Children
	cursor.currentNode.Nature = ASTN_NOTHING
	cursor.currentNode.Name = ""
	cursor.currentNode.src = lexer.Token{}
	cursor.currentNode.Children = []*AstNode{}
	_ = finishAstNode(cursor) // close child and point to here
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Name = formerName
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].src = formerSrc
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Nature = formerNature
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Children = formerChildren
	debug(AST_ROLNE_ITEM, "CCRIWJV-end", cursor)
}

func childParseAstRolneItemFinish(cursor *parseCursor) error {
	// only be called from AST_ROLNE_ITEM_NAME, AST_ROLNE_ITEM_TYPE, or AST_ROLNE_ITEM_VALUE
	debug(AST_ROLNE_ITEM, "CPARIF-start", cursor)

	if cursor.currentNode.Kind == AST_ROLNE_ITEM_NAME {
		childCloseRolneItemWithJustValue(cursor) // re-arrange and point to here
	} else {
		_ = finishAstNode(cursor) // just point to here
	}
	debug(AST_ROLNE_ITEM, "CPARIF-end1", cursor)
	_ = finishAstNode(cursor) // point to parent AST_ROLNE
	debug(AST_ROLNE_ITEM, "CPARIF-end2", cursor)
	return childRolneItemDoneReadyForNextItem(cursor) // have the parent handle the new token
}

func parseAstRolneItemMoveNameToChild(cursor *parseCursor) {
	debug(AST_ROLNE_ITEM, "PARIMNTC", cursor)
	formerName := cursor.currentNode.Children[ROLEITEM_NAMECHILD].Name
	formerSrc := cursor.currentNode.Children[ROLEITEM_NAMECHILD].src
	formerNature := cursor.currentNode.Children[ROLEITEM_NAMECHILD].Nature
	formerChildren := cursor.currentNode.Children[ROLEITEM_NAMECHILD].Children
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Name = formerName
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].src = formerSrc
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Nature = formerNature
	cursor.currentNode.Children[ROLEITEM_VALUECHILD].Children = formerChildren
	cursor.currentNode.Children[ROLEITEM_NAMECHILD].Name = ""
	cursor.currentNode.Children[ROLEITEM_NAMECHILD].src = lexer.Token{}
	cursor.currentNode.Children[ROLEITEM_NAMECHILD].Nature = ASTN_NOTHING
	cursor.currentNode.Children[ROLEITEM_NAMECHILD].Children = []*AstNode{}
}
