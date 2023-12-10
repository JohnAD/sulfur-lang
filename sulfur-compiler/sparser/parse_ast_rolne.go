package sparser

import (
	"fmt"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
)

// In this language, a ROLNE can hold multiple roles:
//   As a collection holding data. A stand-alone pair of single curly braces:  { }
//   As the initial content of a class instance. A pair of parens "bound" to an identifier:  Abc(  )
//   As a parameter definition for a method. A pair of single curly braces as consumed by map-block: parameters {  }
//     When used as a parameter definition, a type is required for each item and the default is the value.
//   As the arguments for a method call. A pair of parens "bound" to an identifier. abc( )
//

// TODO: explore the idea of a "serializer" ROLE using {# #} aka
//    sz::Int = _ {#
//      desc = "Size of bling"
//      yaml = "bling_size"
//      json = "blingSize"
//      required = true
//    #}

func parseAstRolne(cursor *parseCursor, token lexer.Token) error {
	debug("PAR", cursor)
	switch token.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if token.Content == "," {
			return nil // ignore commas between items
		}
	//case lexer.TT_OPEN_SYMBOL:
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstRolne(cursor, token)
	//case lexer.TT_OPEN_BIND_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//case lexer.TT_BINDING_SYMBOL:
	//	return binderHandling(cursor, token)
	case lexer.TT_IDENT:
		return parseAstRolneItemStart(cursor, token)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemStart(cursor, token)
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
	debug("PARS", cursor)
	selfPtr := cursor.currentNode
	selfPtr.Kind = AST_ROLNE
	selfPtr.Nature = nature
	selfPtr.Name = token.Content
	return nil
}

func parseAstRolneStartChild(cursor *parseCursor, token lexer.Token, nature AstNodeNature, bound bool) error {
	debug("PARSC", cursor)
	createAndBecomeChild(cursor, AST_ROLNE, nature, token.Content, bound)
	return nil
}

func finishAstRolne(cursor *parseCursor, token lexer.Token) error {
	debug("FAR", cursor)
	starting := cursor.currentNode.Name
	ending := token.Content
	if helpers.OperatorsMatch(starting, ending) {
		err := finishAstNode(cursor)
		debug("FAR-end", cursor)
		return err
	}
	return fmt.Errorf("[PARSE_ROLNE_FAR] unable to match '%s' with '%s' on line %d column %d", starting, ending, token.SourceLine, token.SourceOffset)
}
