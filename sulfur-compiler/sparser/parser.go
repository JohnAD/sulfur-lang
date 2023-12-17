package sparser

import (
	"fmt"
	"sulfur-compiler/context"
	"sulfur-compiler/lexer"
)

type parseCursor struct {
	depth       int
	src         lexer.Token
	currentNode *AstNode
	pointerPath []*AstNode
}

func addChildAstPointer(cursor *parseCursor, ast *AstNode) {
	cursor.currentNode.Children = append(cursor.currentNode.Children, ast)
}

func addChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature, name string) {
	newNode := AstNode{Kind: ant, Nature: nature, Name: name, src: cursor.src}
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

func createAndBecomeChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature) {
	newNode := AstNode{Kind: ant, Nature: nature, Name: cursor.src.Content, src: cursor.src}
	addChildAstPointer(cursor, &newNode)
	_ = moveToLastChild(cursor) // cannot error out because we just added a child
}

func createAndBecomeEmptyChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature) {
	newNode := AstNode{Kind: ant, Nature: nature, Name: "", src: lexer.Token{}}
	addChildAstPointer(cursor, &newNode)
	_ = moveToLastChild(cursor) // cannot error out because we just added a child
}

func createIndependentChildAndPoint(cursor *parseCursor, nature AstNodeNature, name string) {
	newNode := AstNode{Kind: cursor.currentNode.Kind, Nature: nature, Name: name, src: cursor.src}
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

func setChild(cursor *parseCursor, index int, ant AstNodeType, nature AstNodeNature) error {
	lastIndex := len(cursor.currentNode.Children) - 1
	if index > lastIndex {
		return fmt.Errorf("[PARSER SC} setting a child (%d) that does not exist", index)
	}
	childPtr := cursor.currentNode.Children[index]
	childPtr.Kind = ant
	childPtr.Nature = nature
	childPtr.Name = cursor.src.Content
	childPtr.src = cursor.src
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

func openSymbolHandlingForNewChild(cursor *parseCursor) error {
	debugGeneric("OSHFNC", cursor)
	if cursor.src.Content == "{" {
		return parseAstRolneStartChild(cursor, ASTN_NULL)
	}
	if cursor.src.Content == "(" {
		return parseAstRolneStartChild(cursor, ASTN_NULL)
	}
	if cursor.src.Content == "{{" {
		return parseAstMapBlockStart(cursor)
	}
	if cursor.src.Content == "[[" {
		return parseAstBlockStartChild(cursor, ASTN_GROUPING)
	}
	return fmt.Errorf("[PARSE_GENERIC_OSHFNC] unable to determine what '%s' is on line %d column %d", cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset)
}
func openSymbolHandlingInPlace(cursor *parseCursor) error {
	debugGeneric("OSHIP", cursor)
	if cursor.src.Content == "{" {
		return parseAstRolneStart(cursor, ASTN_NULL)
	}
	if cursor.src.Content == "[[" {
		return parseAstBlockStart(cursor, ASTN_NULL)
	}
	if cursor.src.Content == "(" {
		return parseAstRolneStart(cursor, ASTN_NULL)
	}
	//if token.Content == "{{" {
	//	return parseAstMapBlockStart(cursor)
	//}
	return fmt.Errorf("[PARSE_GENERIC_OSHIP] unable to determine what '%s' is on line %d column %d", cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset)
}

func openSymbolHandlingForLastChild(cursor *parseCursor) error {
	lastIndex := len(cursor.currentNode.Children) - 1
	if lastIndex == -1 {
		return fmt.Errorf("[PARSER_GENERIC_OSHFLC} using a child when the list is empty")
	}
	err := gotoChild(cursor, lastIndex)
	if err != nil {
		return fmt.Errorf("[PARSE_GENERIC_OSHFLC] error: %v", err)
	}
	if cursor.src.Content == "(" {
		return parseAstRolneStartChild(cursor, ASTN_ROLNE_ARGUMENTS)
	}
	return fmt.Errorf("[PARSE_GENERIC_OSHFLC] unable to determine what '%s' is on line %d column %d", cursor.src.Content, cursor.src.SourceLine, cursor.src.SourceOffset)
}

func becomeLastChildMakePreviousChildAChildThenBecomeChild(cursor *parseCursor, nature AstNodeNature, name string) error {
	// before:
	//            a[ b[ d, e, f ] ]            where "b" is the current location
	// after calling with g:
	//            a[ b[ d, e, g[ f ] ] ]       where "g" is the current location
	//
	// used for infix style things. So that a\b becomes `\`[a, b]
	debugGeneric("BLCMPCACTBC", cursor)
	err, previousLastChild := popChild(cursor)
	if err == nil {
		addChild(cursor, previousLastChild.Kind, nature, name)
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
	currentSrc := cursor.currentNode.src
	addChild(cursor, currentKind, currentNature, currentName)
	cursor.currentNode.Kind = ant
	cursor.currentNode.Nature = nature
	cursor.currentNode.Name = name
	cursor.currentNode.src = currentSrc
	return nil
}

func parse(cursor *parseCursor, token lexer.Token) error {
	debugNext(cursor)
	cursor.src = token
	switch cursor.currentNode.Kind {
	case AST_ROOT:
		return parseAstRoot(cursor)
	case AST_ROUTINE:
	case AST_STATEMENT:
		return parseAstStatement(cursor)
	case AST_STATEMENT_ITEM:
		return parseAstStatementItem(cursor)
	case AST_ORDERED_BINDING:
		return parseAstOrderedBinding(cursor)
	case AST_ORDERED_BINDING_CHILD:
		return parseAstOrderedBindingChild(cursor)
	case AST_EXPRESSION:
	case AST_LITERAL:
	case AST_IDENTIFIER:
	case AST_ROLNE:
		return parseAstRolne(cursor)
	case AST_ROLNE_ITEM:
		return parseAstRolneItem(cursor)
	case AST_ROLNE_ITEM_NAME:
		return parseAstRolneItemName(cursor)
	case AST_ROLNE_ITEM_TYPE:
		return parseAstRolneItemType(cursor)
	case AST_ROLNE_ITEM_VALUE:
		return parseAstRolneItemValue(cursor)
	case AST_BLOCK:
		return parseAstBlock(cursor)
	case AST_MAPBLOCK:
		return parseAstMapBlock(cursor)
	case AST_MAPBLOCK_ITEM:
		return parseAstMapBlockItem(cursor)
	case AST_ERROR:
	default:
		return fmt.Errorf("unhandled parse of %v on token %v", cursor.currentNode.Kind, token)
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
	//fmt.Println("CURSOR: ")
	//fmt.Printf("%v", cursor)
	return err, root
}
