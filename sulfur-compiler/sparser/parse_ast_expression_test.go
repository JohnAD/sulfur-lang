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
			kind: ROOT
			nature: _
			name: ""
			children:
				- kind: STATEMENT
				  nature: '?'
				  name: ""
				  children:
					- kind: S-ITEM
					  nature: IDENT
					  name: x
					- kind: S-ITEM
					  nature: INFIX-OP
					  name: =
					- kind: EXPR
					  nature: BINDING
					  name: (
					  children:
						- kind: E-ITEM
						  nature: NUMSTR
						  name: "1"
		`),
	}, {
		"expression with op",
		helpers.Dedent(`
			x = ( 1 + 2 )
		`), helpers.Dedent(`
			kind: ROOT
			nature: _
			name: ""
			children:
				- kind: STATEMENT
				  nature: '?'
				  name: ""
				  children:
					- kind: S-ITEM
					  nature: IDENT
					  name: x
					- kind: S-ITEM
					  nature: INFIX-OP
					  name: =
					- kind: EXPR
					  nature: BINDING
					  name: (
					  children:
						- kind: E-ITEM
						  nature: NUMSTR
						  name: "1"
						- kind: E-ITEM
						  nature: INFIX-OP
						  name: +
						- kind: E-ITEM
						  nature: NUMSTR
						  name: "2"
		`),
	}, {
		"expression with lightly extended op",
		helpers.Dedent(`
			x = ( 1 + 2 == 3 )
		`), helpers.Dedent(`
			kind: ROOT
			nature: _
			name: ""
			children:
				- kind: STATEMENT
				  nature: '?'
				  name: ""
				  children:
					- kind: S-ITEM
					  nature: IDENT
					  name: x
					- kind: S-ITEM
					  nature: INFIX-OP
					  name: =
					- kind: EXPR
					  nature: BINDING
					  name: (
					  children:
						- kind: E-ITEM
						  nature: NUMSTR
						  name: "1"
						- kind: E-ITEM
						  nature: INFIX-OP
						  name: +
						- kind: E-ITEM
						  nature: NUMSTR
						  name: "2"
						- kind: E-ITEM
						  nature: INFIX-OP
						  name: ==
						- kind: E-ITEM
						  nature: NUMSTR
						  name: "3"
		`),
	},
}

func TestExpressionParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	cc.WriteToDisk = false
	cc.ShortYAML = true
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
