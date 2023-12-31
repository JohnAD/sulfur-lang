package sparser

import (
	"fmt"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
)

// In this language, a ROLNE can hold multiple roles:
//   As a collection holding data. A stand-alone pair of single curly braces:  { }
//   ?? As the initial content of a class instance. A pair of parens "bound" to an identifier:  Abc(  )
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

func parseAstRolne(cursor *parseCursor) error {
	debug(cursor, "MAIN")
	switch cursor.src.TokenType {
	case lexer.TT_STANDING_SYMBOL:
		if cursor.src.Content == "," {
			return nil // ignore commas between items
		}
	//case lexer.TT_OPEN_SYMBOL:
	case lexer.TT_CLOSE_SYMBOL:
		return finishAstRolne(cursor)
	case lexer.TT_IDENT:
		return parseAstRolneItemStart(cursor)
	case lexer.TT_STR_LIT:
		return parseAstRolneItemStart(cursor)
	case lexer.TT_NUMSTR_LIT:
		return parseAstRolneItemStart(cursor)
	case lexer.TT_LINE:
		return nil // just ignore EOL when in a ROLNE (but not inside a ROLNE ITEM)
	}
	return parseError(cursor, "MAIN", "unhandled parse")
}

func parseAstRolneStart(cursor *parseCursor, nature AstNodeNature) error {
	debug(cursor, "PARS")
	selfPtr := cursor.currentNode
	selfPtr.Kind = AST_ROLNE
	selfPtr.Nature = nature
	selfPtr.Src = cursor.src
	return nil
}

func parseAstRolneStartChild(cursor *parseCursor, nature AstNodeNature) error {
	debug(cursor, "PARSC")
	createAndBecomeChild(cursor, AST_ROLNE, nature)
	return nil
}

func childRolneItemDoneReadyForNextItem(cursor *parseCursor) error {
	debug(cursor, "CDRFNI")
	return parseAstRolne(cursor)
}

func finishAstRolne(cursor *parseCursor) error {
	debug(cursor, "FAR")
	starting := cursor.currentNode.Src.Content
	ending := cursor.src.Content
	if helpers.OperatorsMatch(starting, ending) {
		err := finishAstNode(cursor)
		debug(cursor, "FAR")
		return err
	}
	return parseError(cursor, "FAR", fmt.Sprintf("unable to match '%s' with '%s'", starting, ending))
}
