package sparser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"sulfur-compiler/context"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
	"testing"
)

var parseExpressionTests = []parseTest{
	{
		"simple expression",
		helpers.Dedent(`
			x = ( 1 )
		`), helpers.Dedent(`
			type: ROOT
			nature: _
			name: ""
			children:
				- type: STATEMENT
				  nature: '?'
				  name: ""
				  children:
					- type: S-ITEM
					  nature: IDENT
					  name: x
					  children: []
					- type: S-ITEM
					  nature: INFIX-OP
					  name: =
					  children: []
					- type: EXPR
					  nature: BINDING
					  name: (
					  children:
						- type: E-ITEM
						  nature: NUMSTR
						  name: "1"
						  children: []
		`),
	}, {
		"expression with op",
		helpers.Dedent(`
			x = ( 1 + 2 )
		`), helpers.Dedent(`
			type: ROOT
			nature: _
			name: ""
			children:
				- type: STATEMENT
				  nature: '?'
				  name: ""
				  children:
					- type: S-ITEM
					  nature: IDENT
					  name: x
					  children: []
					- type: S-ITEM
					  nature: INFIX-OP
					  name: =
					  children: []
					- type: EXPR
					  nature: BINDING
					  name: (
					  children:
						- type: E-ITEM
						  nature: NUMSTR
						  name: "1"
						  children: []
						- type: E-ITEM
						  nature: INFIX-OP
						  name: +
						  children: []
						- type: E-ITEM
						  nature: NUMSTR
						  name: "2"
						  children: []
		`),
	}, {
		"expression with lightly extended op",
		helpers.Dedent(`
			x = ( 1 + 2 == 3 )
		`), helpers.Dedent(`
			type: ROOT
			nature: _
			name: ""
			children:
				- type: STATEMENT
				  nature: '?'
				  name: ""
				  children:
					- type: S-ITEM
					  nature: IDENT
					  name: x
					  children: []
					- type: S-ITEM
					  nature: INFIX-OP
					  name: =
					  children: []
					- type: EXPR
					  nature: BINDING
					  name: (
					  children:
						- type: E-ITEM
						  nature: NUMSTR
						  name: "1"
						  children: []
						- type: E-ITEM
						  nature: INFIX-OP
						  name: +
						  children: []
						- type: E-ITEM
						  nature: NUMSTR
						  name: "2"
						  children: []
						- type: E-ITEM
						  nature: INFIX-OP
						  name: ==
						  children: []
						- type: E-ITEM
						  nature: NUMSTR
						  name: "3"
						  children: []
		`),
	},
}

func TestExpressionParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	cc.WriteToDisk = false
	for _, test := range parseExpressionTests {
		err, tokens := lexer.LexString(test.source)
		if err != nil {
			t.Errorf("before parse test even started, failed to lex on test `%s` LEX:\n%v\nERR: %s", test.desc, tokens, err.Error())
			return
		}
		err, ast := ParseTokensToAst(&cc, tokens)
		rootYamlBytes, _ := yaml.Marshal(ast)
		rootYaml := string(rootYamlBytes)
		if err != nil {
			fmt.Println("AST: ")
			fmt.Println(rootYaml)
			fmt.Println()
			t.Errorf(">>>> on test `%s`, error: %v", test.desc, err)
		}
		if rootYaml != test.expectedYaml {
			t.Errorf(">>>> on test `%s`, GOT:\n%s<<<EOF\n\nEXPECTED:\n%s<<<EOF", test.desc, rootYaml, test.expectedYaml)
		}
	}
}
