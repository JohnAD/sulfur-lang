package sparser

import (
	"fmt"
)

const debugParse = false

func debug(file AstNodeType, location string, cursor *parseCursor) {
	if debugParse {
		fmt.Printf("[PARSE %s.%s] %s\n", file, location, cursor.currentNode.Kind)
	}
}

func debugGeneric(location string, cursor *parseCursor) {
	if debugParse {
		fmt.Printf("[PARSE %s] %s\n", location, cursor.currentNode.Kind)
	}
}

func debugNext(cursor *parseCursor) {
	if debugParse {
		fmt.Printf("== %s token '%s' ==\n", cursor.currentNode.Kind, cursor.src.Content)
	}
}

func parseError(cursor *parseCursor, msg string) error {
	return fmt.Errorf("at token `%s` at line %d col %d: %s", cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset, msg)
}
