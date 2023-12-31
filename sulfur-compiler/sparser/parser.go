package sparser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"sulfur-compiler/context"
	"sulfur-compiler/lexer"
)

type parseCursor struct {
	depth       int
	src         lexer.Token
	currentNode *AstNode
	pointerPath []*AstNode
	cc          context.CompilerContext
}

func addChildAstPointer(cursor *parseCursor, ast *AstNode) {
	cursor.currentNode.Children = append(cursor.currentNode.Children, ast)
}

func addChildWithCopyOfCurrent(cursor *parseCursor, newKind AstNodeType) {
	newNode := AstNode{
		Kind:      newKind,
		Nature:    cursor.currentNode.Nature,
		Src:       cursor.currentNode.Src,
		shortYAML: cursor.cc.ShortYAML,
	}
	cursor.currentNode.Children = append(cursor.currentNode.Children, &newNode)
}

func addChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature) {
	newNode := AstNode{Kind: ant, Nature: nature, Src: cursor.src, shortYAML: cursor.cc.ShortYAML}
	addChildAstPointer(cursor, &newNode)
}

func addNullChild(cursor *parseCursor, ant AstNodeType) {
	newNode := AstNode{Kind: ant, Nature: ASTN_NULL, Src: lexer.Token{}, shortYAML: cursor.cc.ShortYAML}
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
	newNode := AstNode{Kind: ant, Nature: nature, Src: cursor.src, shortYAML: cursor.cc.ShortYAML}
	addChildAstPointer(cursor, &newNode)
	_ = moveToLastChild(cursor) // cannot error out because we just added a child
}

func createAndBecomeEmptyChild(cursor *parseCursor, ant AstNodeType, nature AstNodeNature) {
	newNode := AstNode{Kind: ant, Nature: nature, Src: lexer.Token{}, shortYAML: cursor.cc.ShortYAML}
	addChildAstPointer(cursor, &newNode)
	_ = moveToLastChild(cursor) // cannot error out because we just added a child
}

func createIndependentChildAndPoint(cursor *parseCursor, nature AstNodeNature) {
	newNode := AstNode{Kind: cursor.currentNode.Kind, Nature: nature, Src: cursor.src, shortYAML: cursor.cc.ShortYAML}
	addChildAstPointer(cursor, &newNode)
	_ = moveToLastChild(cursor) // cannot error out because we just added a child
}

func applyNameNatureToSelf(cursor *parseCursor, nature AstNodeNature) {
	cursor.currentNode.Src = cursor.src
	cursor.currentNode.Nature = nature
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
	childPtr.Src = cursor.src
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

func openBindSymbolHandlingForNewChild(cursor *parseCursor) error {
	debug(cursor, "OBSHFNC")
	//if cursor.Src.Content == "{" {
	//	return parseAstRolneStartChild(cursor, ASTN_NULL)
	//}
	if cursor.src.Content == "(" {
		return parseAstRolneStartChild(cursor, ASTN_NULL)
	}
	//if cursor.Src.Content == "{{" {
	//	return parseAstMapBlockStart(cursor)
	//}
	//if cursor.Src.Content == "[[" {
	//	return parseAstBlockStartChild(cursor, ASTN_GROUPING)
	//}
	return parseError(cursor, "OBSHFNC", fmt.Sprintf("unable to determine what '%s' is", cursor.src.Content))
}

func openSymbolHandlingForNewChild(cursor *parseCursor) error {
	debug(cursor, "OSHFNC")
	if cursor.src.Content == "{" {
		return parseAstRolneStartChild(cursor, ASTN_NULL)
	}
	if cursor.src.Content == "(" {
		return parseAstExpressionStartChild(cursor)
	}
	if cursor.src.Content == "{{" {
		return parseAstMapBlockStart(cursor)
	}
	if cursor.src.Content == "[[" {
		return parseAstBlockStartChild(cursor, ASTN_GROUPING)
	}
	return parseError(cursor, "OSHFNC", "unable to determine")
}

func openSymbolHandlingInPlace(cursor *parseCursor) error {
	debug(cursor, "OSHIP")
	if cursor.src.Content == "{" {
		return parseAstRolneStart(cursor, ASTN_NULL)
	}
	if cursor.src.Content == "[[" {
		return parseAstBlockStart(cursor, ASTN_NULL)
	}
	if cursor.src.Content == "(" {
		return parseAstRolneStart(cursor, ASTN_NULL)
	}
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
		return parseAstExpression(cursor)
	case AST_EXPRESSION_ITEM:
		return parseAstExpressionItem(cursor)
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
	//case AST_INFIX:
	//	return parseAstInfix(cursor)
	//case AST_INFIX_RIGHT:
	//	return parseAstInfixRight(cursor)
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
		cc:          *cc,
	}
	root := AstNode{shortYAML: cursor.cc.ShortYAML}
	cursor.currentNode = &root
	for _, token := range tokens {
		err = parse(&cursor, token)
		if err != nil {
			//fmt.Println("CURSOR BEFORE ERROR: ")
			//fmt.Printf("%v\n", cursor)
			return err, root
		}
	}
	yamlData, err := yaml.Marshal(&root)
	if err != nil {
		return err, root
	}
	if cc.WriteToDisk {
		fileId := 0 // TODO: handling file numbers later
		outPath := context.GetParseResultPath(cc)
		yamlFilePath := outPath + "/" + fmt.Sprintf("file-%04d.ast.yaml", fileId)
		err = os.WriteFile(yamlFilePath, yamlData, 0644)
	}
	return err, root
}
