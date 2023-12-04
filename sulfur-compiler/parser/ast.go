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
	AST_TYPE
	AST_VALUE
	AST_TARGET
	AST_BLOCK
	AST_MAPBLOCK
	AST_MAPBLOCK_ITEM
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
	case AST_TYPE:
		return "TYPE"
	case AST_VALUE:
		return "VALUE"
	case AST_TARGET:
		return "TARGET"
	case AST_INFIX:
		return "INFIX"
	case AST_BLOCK:
		return "BLOCK"
	case AST_MAPBLOCK:
		return "MBLOCK"
	case AST_MAPBLOCK_ITEM:
		return "MB-ITEM"
	case AST_ERROR:
		return "ERROR"
	default:
		return fmt.Sprintf("%d", int(ant))
	}
}

func (ant AstNodeType) MarshalYAML() (interface{}, error) {
	return ant.String(), nil
}

type AstNodeNature int

const (
	ASTN_NOTHING AstNodeNature = iota
	ASTN_NULL
	ASTN_STATEMENT_ROOT_DECLARATION
	ASTN_STATEMENT_ROOT_FRAMEWORK
	ASTN_STATEMENT_ASSIGN
	ASTN_IDENTIFIER
	ASTN_STR
	ASTN_NUMSTR
	ASTN_INFIX_OPERATOR
	ASTN_META_BINDING
	ASTN_GROUPING
	ASTN_KEYWORD
)

func (ann AstNodeNature) String() string {
	switch ann {
	case ASTN_NOTHING:
		return "_"
	case ASTN_NULL:
		return "?"
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
	case ASTN_GROUPING:
		return "GROUP"
	case ASTN_KEYWORD:
		return "KEYWORD"
	default:
		return fmt.Sprintf("%d", int(ann))
	}
}

func (ann AstNodeNature) MarshalYAML() (interface{}, error) {
	return ann.String(), nil
}

type AstNode struct {
	Kind     AstNodeType   `yaml:"type"`
	Nature   AstNodeNature `yaml:"nature"`
	Name     string        `yaml:"name"`
	Children []*AstNode    `yaml:"children"`
}

func (an AstNode) String() string {
	if len(an.Children) > 0 {
		return fmt.Sprintf("AST(%s.%s.`%s` %v)", an.Kind, an.Nature, an.Name, an.Children)
	}
	return fmt.Sprintf("AST(%s.%s.`%s`)", an.Kind, an.Nature, an.Name)
}
