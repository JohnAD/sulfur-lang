package sparser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"sulfur-compiler/context"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
	"testing"
)

var parseRolneDataTests = []parseTest{
	parseTest{
		"simple vertical",
		helpers.Dedent(`
			let t = { 
              "a" = "first"
              "second"
              "c" = "third"
            }
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
					  name: let
					  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: t
					  children: []
					- type: S-ITEM
					  nature: INFIX-OP
					  name: =
					  children: []
					- type: ROLNE
					  nature: '?'
					  name: '{'
					  children:
						- type: R-ITEM
						  nature: _
						  name: ""
						  children:
							- type: R-I-NAME
							  nature: STR-LIT
							  name: a
							  children: []
							- type: R-I-TYPE
							  nature: '?'
							  name: ""
							  children: []
							- type: R-I-VALUE
							  nature: STR-LIT
							  name: first
							  children: []
						- type: R-ITEM
						  nature: _
						  name: ""
						  children:
							- type: R-I-NAME
							  nature: _
							  name: ""
							  children: []
							- type: R-I-TYPE
							  nature: '?'
							  name: ""
							  children: []
							- type: R-I-VALUE
							  nature: STR-LIT
							  name: second
							  children: []
						- type: R-ITEM
						  nature: _
						  name: ""
						  children:
							- type: R-I-NAME
							  nature: STR-LIT
							  name: c
							  children: []
							- type: R-I-TYPE
							  nature: '?'
							  name: ""
							  children: []
							- type: R-I-VALUE
							  nature: STR-LIT
							  name: third
							  children: []
		`),
	},
	parseTest{
		"simple horizontal", helpers.Dedent(`
			let t = { "a" = "first" , "second" , "c" = "third" }
		`),
		helpers.Dedent(`
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
					  name: let
					  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: t
					  children: []
					- type: S-ITEM
					  nature: INFIX-OP
					  name: =
					  children: []
					- type: ROLNE
					  nature: '?'
					  name: '{'
					  children:
						- type: R-ITEM
						  nature: _
						  name: ""
						  children:
							- type: R-I-NAME
							  nature: STR-LIT
							  name: a
							  children: []
							- type: R-I-TYPE
							  nature: '?'
							  name: ""
							  children: []
							- type: R-I-VALUE
							  nature: STR-LIT
							  name: first
							  children: []
						- type: R-ITEM
						  nature: _
						  name: ""
						  children:
							- type: R-I-NAME
							  nature: _
							  name: ""
							  children: []
							- type: R-I-TYPE
							  nature: '?'
							  name: ""
							  children: []
							- type: R-I-VALUE
							  nature: STR-LIT
							  name: second
							  children: []
						- type: R-ITEM
						  nature: _
						  name: ""
						  children:
							- type: R-I-NAME
							  nature: STR-LIT
							  name: c
							  children: []
							- type: R-I-TYPE
							  nature: '?'
							  name: ""
							  children: []
							- type: R-I-VALUE
							  nature: STR-LIT
							  name: third
							  children: []
		`),
	},
}

func TestRolneDataParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	for _, test := range parseRolneDataTests {
		err, tokens := lexer.LexString(test.source)
		if err != nil {
			t.Errorf("before parse test even started, failed to lex on test %s", test.desc)
		}
		err, ast := ParseTokensToAst(&cc, tokens)
		rootYamlBytes, _ := yaml.Marshal(ast)
		rootYaml := string(rootYamlBytes)
		if err != nil {
			fmt.Println("AST: ")
			fmt.Println(rootYaml)
			fmt.Println()
			t.Errorf("on test %s, error %v", test.desc, err)
		}
		if rootYaml != test.expectedYaml {
			t.Errorf("on test %s, GOT:\n%s<<<EOF\n\nEXPECTED:\n%s<<<EOF", test.desc, rootYaml, test.expectedYaml)
		}
	}
}
