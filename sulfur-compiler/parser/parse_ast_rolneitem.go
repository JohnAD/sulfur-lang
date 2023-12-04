package parser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

type RoleItemState int // for answer the question: where are we?
const (
	RIS_NAME RoleItemState = iota
	RIS_TYPE
	RIS_VALUE
)

const ROLEITEM_TYPECHILD = 0
const ROLEITEM_VALUECHILD = 1

// every AST_ROLNE_ITEM should have EXACTLY THREE ITEMS: name type value

func parseAstRolneItem(cursor *parseCursor, token lexer.Token) error {
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if token.Content == "=" {
			return parseAstRolneItemAssignment(cursor, token)
		}
	//case lexer.TT_OPEN_SYMBOL:
	//case lexer.TT_CLOSE_SYMBOL:
	//case lexer.TT_OPEN_BIND_SYMBOL:
	case lexer.TT_BINDING_SYMBOL:
		if token.Content == "::" {
			return parseAstRolneItemTypeStart(cursor, token)
		}
		return parseAstRolneItemOtherBinding(cursor, token)
	case lexer.TT_IDENT:
		return parseAstRolneItemNewPart(cursor, token, ASTN_IDENTIFIER)
	//case lexer.TT_STR_LIT:
	//case lexer.TT_NUMSTR_LIT:
	//case lexer.TT_SYNTAX_ERROR:
	//case lexer.TT_COMMENT:
	//case lexer.TT_WHITESPACE: TODO: remove whitespace TT
	//	return nil
	case lexer.TT_LINE:
		return finishAstRolneItem(cursor)
	}
	return fmt.Errorf("unhandled AST_ROLNE_ITEM parse of %v", token)
}

func parseAstRolneItemNewPart(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	ris := getRolneItemState(cursor)
	switch ris {
	case RIS_TYPE:
		typeAst := rolneItemPointAtType(cursor)
		typeAst.Nature = nature
		typeAst.Name = token.Content
		return nil
	case RIS_VALUE:
		valueAst := rolneItemPointAtValue(cursor)
		valueAst.Nature = nature
		valueAst.Name = token.Content
		return nil
	}
	return fmt.Errorf("[PARSE_ROLNEITEM_PARINV] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}

func parseAstRolneItemOtherBinding(cursor *parseCursor, token lexer.Token) error {
	ris := getRolneItemState(cursor)
	switch ris {
	case RIS_VALUE:
		valueAst := rolneItemPointAtValue(cursor)
		if valueAst.Nature != ASTN_NOTHING { // you can't do <nothing>::Int (not a valid value)
			err := gotoChild(cursor, ROLEITEM_VALUECHILD)
			if err != nil {
				return err
			}
			return binderHandlingInPlace(cursor, token) // changes AST state to AST_ORDERED_BINDING
		}
	}
	return fmt.Errorf("[PARSE_ROLNEITEM_PARIOB] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}

func finishAstRolneItem(cursor *parseCursor) error {
	//fmt.Printf("PING at %v\n", cursor.currentNode)
	if getRolneItemState(cursor) != RIS_VALUE {
		parseAstRolneItemHandleNoName(cursor) // this does a swap of name and value
	}
	err := finishAstNode(cursor)
	//fmt.Printf("   now at %v\n", cursor.currentNode)
	return err
}

func parseAstRolneItemStart(cursor *parseCursor, token lexer.Token, nature AstNodeNature) error {
	createAndBecomeChild(cursor, AST_ROLNE_ITEM, nature, token.Content) // create/become R-ITEM
	addChild(cursor, AST_TYPE, ASTN_NULL, "")                           // add yet-unknown type
	addChild(cursor, AST_VALUE, ASTN_NULL, "")                          // add yet-unknown value
	return nil
}

func parseAstRolneItemAssignment(cursor *parseCursor, token lexer.Token) error {
	rolneItemPointAtValue(cursor).Nature = ASTN_NOTHING // by changing from null(unknown) to nothing, this parser now knows we are working on the value not the name
	return nil
}

func parseAstRolneItemTypeStart(cursor *parseCursor, token lexer.Token) error {
	ris := getRolneItemState(cursor)
	switch ris {
	case RIS_NAME:
		rolneItemPointAtType(cursor).Nature = ASTN_NOTHING // changing from null(unknown) to nothing
		return nil
	case RIS_TYPE:
		// already working on the type, so bind it like normal
		typeAst := rolneItemPointAtType(cursor)
		if typeAst.Nature != ASTN_NOTHING { // you can't do <nothing>::Int (not a valid type)
			err := gotoChild(cursor, ROLEITEM_TYPECHILD)
			if err != nil {
				return err
			}
			return binderHandlingInPlace(cursor, token) // changes AST state to AST_ORDERED_BINDING
		}
	}
	return fmt.Errorf("[PARSE_ROLNEITEM_PARITS] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}

func parseAstRolneItemHandleNoName(cursor *parseCursor) {
	formerName := cursor.currentNode.Name
	cursor.currentNode.Nature = ASTN_NOTHING
	cursor.currentNode.Name = ""
	rolneItemPointAtValue(cursor).Nature = ASTN_IDENTIFIER
	rolneItemPointAtValue(cursor).Name = formerName
}

func rolneItemPointAtType(cursor *parseCursor) *AstNode {
	return cursor.currentNode.Children[ROLEITEM_TYPECHILD]
}

func rolneItemPointAtValue(cursor *parseCursor) *AstNode {
	return cursor.currentNode.Children[ROLEITEM_VALUECHILD]
}

func getRolneItemState(cursor *parseCursor) RoleItemState {
	if rolneItemPointAtValue(cursor).Nature != ASTN_NULL {
		return RIS_VALUE
	}
	if rolneItemPointAtType(cursor).Nature != ASTN_NULL {
		return RIS_TYPE
	}
	return RIS_NAME
}
