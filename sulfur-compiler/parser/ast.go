package parser

import (
	"fmt"
)

type AstNodeType int

const (
	AST_ROOT AstNodeType = iota
	AST_ROUTINE
	AST_STATEMENT
	AST_STATEMENT_ITEM
	AST_ORDERED_BINDING
	AST_EXPRESSION
	AST_LITERAL
	AST_IDENTIFIER
	AST_ROLNE
	AST_ROLNE_ITEM
	AST_INFIX
	AST_ERROR
)

func (ant AstNodeType) String() string {
	switch ant {
	case AST_ROOT:
		return "ROOT"
	case AST_ROUTINE:
		return "ROUTINE"
	case AST_STATEMENT:
		return "STATEMENT"
	case AST_STATEMENT_ITEM:
		return "S-ITEM"
	case AST_ORDERED_BINDING:
		return "BINDING"
	case AST_EXPRESSION:
		return "EXPRESSION"
	case AST_LITERAL:
		return "LITERAL"
	case AST_IDENTIFIER:
		return "IDENTIFIER"
	case AST_ROLNE:
		return "ROLNE"
	case AST_ROLNE_ITEM:
		return "R-ITEM"
	case AST_INFIX:
		return "INFIX"
	case AST_ERROR:
		return "ERROR"
	default:
		return fmt.Sprintf("%d", int(ant))
	}
}

type AstNodeNature int

const (
	ASTN_ALL_EMPTY AstNodeNature = iota
	ASTN_STATEMENT_ROOT_DECLARATION
	ASTN_STATEMENT_ROOT_FRAMEWORK
	ASTN_STATEMENT_ASSIGN
	ASTN_IDENTIFIER
	ASTN_STR
	ASTN_NUMSTR
	ASTN_INFIX_OPERATOR
	ASTN_META_BINDING
)

func (ann AstNodeNature) String() string {
	switch ann {
	case ASTN_ALL_EMPTY:
		return "_"
	case ASTN_STATEMENT_ROOT_DECLARATION:
		return "ROOT-DECL"
	case ASTN_STATEMENT_ROOT_FRAMEWORK:
		return "ROOT-FW"
	case ASTN_STATEMENT_ASSIGN:
		return "ASSIGN"
	case ASTN_IDENTIFIER:
		return "IDENT"
	case ASTN_STR:
		return "STR"
	case ASTN_NUMSTR:
		return "NUMSTR"
	case ASTN_INFIX_OPERATOR:
		return "INFIX-OP"
	case ASTN_META_BINDING:
		return "BINDING"
	default:
		return fmt.Sprintf("%d", int(ann))
	}
}

type AstNode struct {
	kind     AstNodeType
	nature   AstNodeNature
	name     string
	children []*AstNode
	// items    map[string]AstNode
}

func (an AstNode) String() string {
	if len(an.children) > 0 {
		return fmt.Sprintf("AST(%s.%s.`%s` %v)", an.kind, an.nature, an.name, an.children)
	}
	return fmt.Sprintf("AST(%s.%s.`%s`)", an.kind, an.nature, an.name)
}
