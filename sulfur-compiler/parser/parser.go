package parser

import (
	"fmt"
	"sulfur-compiler/context"
	"sulfur-compiler/lexer"
)

type parseCursor struct {
	depth       int
	currentNode *AstNode
	pointerPath []*AstNode
}

func addChildAstPointer(cursor *parseCursor, ast *AstNode) {
	cursor.currentNode.Children = append(cursor.currentNode.Children, ast)
}

func addChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature, name string) {
	newNode := AstNode{Kind: ant, Nature: nature, Name: name}
	addChildAstPointer(cursor, &newNode)
}

func moveToLastChild(cursor *parseCursor) error {
	var err error
	childLen := len(cursor.currentNode.Children)
	if childLen == 0 {
		err = fmt.Errorf("[PARSER VTLC} attempting access a child node when array is empty")
	} else {
		ast := cursor.currentNode.Children[childLen-1]
		cursor.pointerPath = append(cursor.pointerPath, cursor.currentNode)
		cursor.currentNode = ast
		cursor.depth += 1
	}
	return err
}

func createAndBecomeChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature, name string) {
	newNode := AstNode{Kind: ant, Nature: nature, Name: name}
	addChildAstPointer(cursor, &newNode)
	_ = moveToLastChild(cursor) // cannot error out because we just added a child
}

func popChild(cursor *parseCursor) (error, *AstNode) {
	// remove last child and return it
	var err error
	var ast *AstNode
	childLen := len(cursor.currentNode.Children)
	if childLen == 0 {
		err = fmt.Errorf("[PARSER PC} attempting to remove a child node when array is empty")
	} else {
		ast = cursor.currentNode.Children[childLen-1]
		cursor.currentNode.Children = cursor.currentNode.Children[:childLen-1]
	}
	return err, ast
}

func gotoChild(cursor *parseCursor, index int) error {
	// start pointing to one of the already-existing children
	lastIndex := len(cursor.currentNode.Children) - 1
	if index > lastIndex {
		return fmt.Errorf("[PARSER GC} moving to a child (%d) that does not exist", index)
	}
	cursor.pointerPath = append(cursor.pointerPath, cursor.currentNode)
	cursor.depth += 1
	cursor.currentNode = cursor.currentNode.Children[index]
	return nil
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

func openSymbolHandling(cursor *parseCursor, token lexer.Token) error {
	if token.Content == "{" {
		return parseAstRolneStartChild(cursor, token)
	}
	if token.Content == "{{" {
		return parseAstMapBlockStart(cursor, token)
	}
	return fmt.Errorf("[PARSE_GENERIC_OSH] unable to determine what '%s' is on line %d column %d", token.Content, token.SourceLine, token.SourceOffset)
}

func becomeLastChildMakePreviousChildAChildThenBecomeChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature, name string) error {
	// before:
	//            a[ b[ d, e, f ] ]            where "b" is the current location
	// after calling with g:
	//            a[ b[ d, e, g[ f ] ] ]       where "g" is the current location
	//
	// used for infix style things. So that a\b becomes `\`[a, b]
	err, previousLastChild := popChild(cursor)
	if err == nil {
		addChild(cursor, ant, nature, name)
		_ = moveToLastChild(cursor) // cannot error out because we just added a child
		addChildAstPointer(cursor, previousLastChild)
	}
	return err
}

func swapSelfMakingPreviousAChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature, name string) error {
	// before:
	//            a[ b[] ]            where "b" is the current location
	// after calling with g
	//            a[ g[ b ] ]         where "g" is the current location
	currentKind := cursor.currentNode.Kind
	currentNature := cursor.currentNode.Nature
	currentName := cursor.currentNode.Name
	addChild(cursor, currentKind, currentNature, currentName)
	cursor.currentNode.Kind = ant
	cursor.currentNode.Nature = nature
	cursor.currentNode.Name = name
	return nil
}

func parse(cursor *parseCursor, token lexer.Token) error {
	switch cursor.currentNode.Kind {
	case AST_ROOT:
		return parseAstRoot(cursor, token)
	case AST_ROUTINE:
	case AST_STATEMENT:
		return parseAstStatement(cursor, token)
	case AST_ORDERED_BINDING:
		return parseAstOrderedBinding(cursor, token)
	case AST_EXPRESSION:
	case AST_LITERAL:
	case AST_IDENTIFIER:
	case AST_ROLNE:
		return parseAstRolne(cursor, token)
	case AST_ROLNE_ITEM:
		return parseAstRolneItem(cursor, token)
	case AST_BLOCK:
	case AST_MAPBLOCK:
		return parseAstMapBlock(cursor, token)
	case AST_MAPBLOCK_ITEM:
		return parseAstMapBlockItem(cursor, token)
	case AST_ERROR:
	default:
		return fmt.Errorf("unhandled parse of %v", cursor.currentNode.Kind)
	}
	return nil
}

func ParseTokensToAst(cc *context.CompilerContext, tokens []lexer.Token) (error, AstNode) {
	var err error = nil
	cursor := parseCursor{
		depth:       0,
		pointerPath: []*AstNode{},
	}
	root := AstNode{}
	cursor.currentNode = &root
	for _, token := range tokens {
		err = parse(&cursor, token)
		if err != nil {
			//fmt.Println("CURSOR BEFORE ERROR: ")
			//fmt.Printf("%v\n", cursor)
			return err, root
		}
	}
	fmt.Println("CURSOR: ")
	fmt.Printf("%v", cursor)
	return err, root
}
