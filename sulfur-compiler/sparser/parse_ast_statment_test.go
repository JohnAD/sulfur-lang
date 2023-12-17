package sparser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"sulfur-compiler/context"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
	"testing"
)

type parseTest struct {
	desc         string
	source       string
	expectedYaml string
}

var parseTests = []parseTest{
	parseTest{
		"simple five parts", helpers.Dedent(`
			) a b c d e
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
					  name: a
					  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: b
					  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: c
					  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: d
					  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: e
					  children: []
    `)},
	parseTest{
		"simple parameterized call", helpers.Dedent(`
			abc( d e )
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
					  name: abc
					  children:
						- type: ROLNE
						  nature: ARGS
						  name: (
						  children:
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
								  nature: IDENT
								  name: d
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
								  nature: IDENT
								  name: e
								  children: []
    `)},
}

func TestStatementParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	for _, test := range parseTests {
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
