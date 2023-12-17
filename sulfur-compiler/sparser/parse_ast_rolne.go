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
	debug(AST_ROLNE, "MAIN", cursor)
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
	case lexer.TT_LINE:
		return nil // just ignore EOL when in a ROLNE (but not inside a ROLNE ITEM)
	}
	return fmt.Errorf("unhandled AST_ROLNE parse of %v", cursor.src)
}

func parseAstRolneStart(cursor *parseCursor, nature AstNodeNature) error {
	debug(AST_ROLNE, "PARS", cursor)
	selfPtr := cursor.currentNode
	selfPtr.Kind = AST_ROLNE
	selfPtr.Nature = nature
	selfPtr.Name = cursor.src.Content
	selfPtr.src = cursor.src
	return nil
}

func parseAstRolneStartChild(cursor *parseCursor, nature AstNodeNature) error {
	debug(AST_ROLNE, "PARSC", cursor)
	createAndBecomeChild(cursor, AST_ROLNE, nature)
	return nil
}

func childRolneItemDoneReadyForNextItem(cursor *parseCursor) error {
	debug(AST_ROLNE, "CDRFNI", cursor)
	return parseAstRolne(cursor)
}

func finishAstRolne(cursor *parseCursor) error {
	debug(AST_ROLNE, "FAR", cursor)
	starting := cursor.currentNode.Name
	ending := cursor.src.Content
	if helpers.OperatorsMatch(starting, ending) {
		err := finishAstNode(cursor)
		debug(AST_ROLNE, "FAR-end", cursor)
		return err
	}
	return fmt.Errorf("[PARSE_ROLNE_FAR] unable to match '%s' with '%s' on line %d column %d", starting, ending, cursor.src.SourceLine, cursor.src.SourceOffset)
}
