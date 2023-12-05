package sparser

import "sulfur-compiler/lexer"

// a "statement-item" aka S-ITEM is a psuedo class representing the children of a statement
//
// if sparser every reaches the AST_STATEMENT_ITEM state, that probably means some kind of "binding" operation
// just completed. The correct behavior is to "finish" and pass the token to the statement sparser to continue
// the parsing.

func parseAstStatementItem(cursor *parseCursor, token lexer.Token) error {
	err := finishAstNode(cursor)
	if err != nil {
		return err
	}
	return parseAstStatement(cursor, token)
}
