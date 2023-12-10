package sparser

import (
	"fmt"
	"sulfur-compiler/lexer"
)

const debugParse = false

func debug(location string, cursor *parseCursor) {
	if debugParse {
		fmt.Printf("[PARSE %s] %s\n", location, cursor.currentNode.Kind)
	}
}

func debugNext(cursor *parseCursor, token lexer.Token) {
	if debugParse {
		fmt.Printf("== %s token '%s' ==\n", cursor.currentNode.Kind, token.Content)
	}
}
