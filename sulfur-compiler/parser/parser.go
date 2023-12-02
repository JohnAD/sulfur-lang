package parser

import (
	"fmt"
	"sulfur-compiler/context"
	"sulfur-compiler/lexer"
)

type parseCursor struct {
	root        AstNode
	depth       int
	currentNode *AstNode
	pointerPath []*AstNode
}

func createAndBecomeChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature, name string) {
	newNode := AstNode{kind: ant, nature: nature, name: name}
	cursor.currentNode.children = append(cursor.currentNode.children, &newNode)
	cursor.pointerPath = append(cursor.pointerPath, cursor.currentNode)
	cursor.currentNode = &newNode
	cursor.depth += 1
}

func addChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature, name string) {
	newNode := AstNode{kind: ant, nature: nature, name: name}
	cursor.currentNode.children = append(cursor.currentNode.children, &newNode)
}

func finishAstNode(cursor *parseCursor) error {
	size := len(cursor.pointerPath)
	if size > 0 {
		cursor.currentNode = cursor.pointerPath[size-1]
		cursor.pointerPath = cursor.pointerPath[:size-1]
		cursor.depth -= 1
	} else {
		return fmt.Errorf("finishAstNode attempt to finish empty path (pointer list)")
	}
	return nil
}

func parse(cursor *parseCursor, token lexer.Token) error {
	switch cursor.currentNode.kind {
	case AST_ROOT:
		return parseAstRoot(cursor, token)
	case AST_ROUTINE:
	case AST_STATEMENT:
		return parseAstStatement(cursor, token)
	case AST_EXPRESSION:
	case AST_LITERAL:
	case AST_IDENTIFIER:
	case AST_ERROR:
	default:
		return fmt.Errorf("unhandled parse of %v", cursor.currentNode.kind)
	}
	return nil
}

func ParseTokensToAst(cc *context.CompilerContext, tokens []lexer.Token) (error, AstNode) {
	var err error = nil
	cursor := parseCursor{
		root:        AstNode{kind: AST_ROOT},
		depth:       0,
		pointerPath: []*AstNode{},
	}
	cursor.currentNode = &cursor.root
	for _, token := range tokens {
		err = parse(&cursor, token)
		if err != nil {
			fmt.Println("CURSOR BEFORE ERROR: ")
			fmt.Printf("%v\n", cursor)
			fmt.Println("ROOT BEFORE ERROR: ")
			fmt.Printf("%v\n", cursor.root)
			return err, cursor.root
		}
	}
	fmt.Println("CURSOR: ")
	fmt.Printf("%v", cursor)
	return err, cursor.root
}
