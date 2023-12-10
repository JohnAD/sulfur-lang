package sparser

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"sulfur-compiler/context"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
	"testing"
)

var parseRolneArgumentTests = []parseTest{
	parseTest{
		"simple arguments",
		helpers.Dedent(`
			let t = abc( "hello" , x )
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
								  nature: STR-LIT
								  name: hello
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
								  name: x
								  children: []
		`),
	},
	parseTest{
		"simple vertical arguments", helpers.Dedent(`
			let t = abc( 
			  "hello"
			  x
			)
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
								  nature: STR-LIT
								  name: hello
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
								  name: x
								  children: []
		`),
	},
	parseTest{
		"arguments in arguments", helpers.Dedent(`
			let t = abc( "hello" , xyz( "def" ) )
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
								  nature: STR-LIT
								  name: hello
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
								  name: xyz
								  children:
									- type: ROLNE
									  nature: '?'
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
											  nature: STR-LIT
											  name: def
											  children: []
		`),
	},
	parseTest{
		"named arguments", helpers.Dedent(`
			let t = abc( bling = "hello" , bam = x )
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
					- type: S-ITEM
					  nature: IDENT
					  name: abc
					  children:
						- type: ROLNE
						  nature: ARGS
						  name: (
						  children:
							- type: R-ITEM
							  nature: IDENT
							  name: bling
							  children:
								- type: TYPE
								  nature: '?'
								  name: ""
								  children: []
								- type: VALUE
								  nature: STR-LIT
								  name: hello
								  children: []
							- type: R-ITEM
							  nature: IDENT
							  name: bam
							  children:
								- type: TYPE
								  nature: '?'
								  name: ""
								  children: []
								- type: VALUE
								  nature: IDENT
								  name: x
								  children: []
		`),
	},
}

func TestRolneArgumentParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	for _, test := range parseRolneArgumentTests {
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
