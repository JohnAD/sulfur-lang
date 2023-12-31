package sparser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"sulfur-compiler/context"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
	"testing"
)

var parseRolneArgumentTests = []parseTest{
	{
		"simple arguments",
		helpers.Dedent(`
				let t = abc( "hello" , x )
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
						  name: let
						- kind: S-ITEM
						  nature: IDENT
						  name: t
						- kind: S-ITEM
						  nature: INFIX-OP
						  name: =
						- kind: S-ITEM
						  nature: IDENT
						  name: abc
						  children:
							- kind: ROLNE
							  nature: ARGS
							  name: (
							  children:
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: _
									  name: ""
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: STR-LIT
									  name: hello
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: _
									  name: ""
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: IDENT
									  name: x
			`),
	}, {
		"simple vertical arguments", helpers.Dedent(`
				let t = abc( 
				  "hello"
				  x
				)
			`),
		helpers.Dedent(`
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
						  name: let
						- kind: S-ITEM
						  nature: IDENT
						  name: t
						- kind: S-ITEM
						  nature: INFIX-OP
						  name: =
						- kind: S-ITEM
						  nature: IDENT
						  name: abc
						  children:
							- kind: ROLNE
							  nature: ARGS
							  name: (
							  children:
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: _
									  name: ""
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: STR-LIT
									  name: hello
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: _
									  name: ""
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: IDENT
									  name: x
			`),
	}, {
		"arguments in arguments", helpers.Dedent(`
				let t = abc( "hello" , xyz( "def" ) )
			`),
		helpers.Dedent(`
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
						  name: let
						- kind: S-ITEM
						  nature: IDENT
						  name: t
						- kind: S-ITEM
						  nature: INFIX-OP
						  name: =
						- kind: S-ITEM
						  nature: IDENT
						  name: abc
						  children:
							- kind: ROLNE
							  nature: ARGS
							  name: (
							  children:
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: _
									  name: ""
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: STR-LIT
									  name: hello
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: _
									  name: ""
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: IDENT
									  name: xyz
									  children:
										- kind: ROLNE
										  nature: '?'
										  name: (
										  children:
											- kind: R-ITEM
											  nature: _
											  name: ""
											  children:
												- kind: R-I-NAME
												  nature: _
												  name: ""
												- kind: R-I-TYPE
												  nature: '?'
												  name: ""
												- kind: R-I-VALUE
												  nature: STR-LIT
												  name: def
			`),
	}, {
		"named arguments", helpers.Dedent(`
				let t = abc( bling = "hello" , bam = x )
			`),
		helpers.Dedent(`
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
						  name: let
						- kind: S-ITEM
						  nature: IDENT
						  name: t
						- kind: S-ITEM
						  nature: INFIX-OP
						  name: =
						- kind: S-ITEM
						  nature: IDENT
						  name: abc
						  children:
							- kind: ROLNE
							  nature: ARGS
							  name: (
							  children:
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: IDENT
									  name: bling
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: STR-LIT
									  name: hello
								- kind: R-ITEM
								  nature: _
								  name: ""
								  children:
									- kind: R-I-NAME
									  nature: IDENT
									  name: bam
									- kind: R-I-TYPE
									  nature: '?'
									  name: ""
									- kind: R-I-VALUE
									  nature: IDENT
									  name: x
			`),
	}, {
		"bound argument", helpers.Dedent(`
				let t = abc( bling.zen )
			`),
		helpers.Dedent(`
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
					  name: let
					- kind: S-ITEM
					  nature: IDENT
					  name: t
					- kind: S-ITEM
					  nature: INFIX-OP
					  name: =
					- kind: S-ITEM
					  nature: IDENT
					  name: abc
					  children:
						- kind: ROLNE
						  nature: ARGS
						  name: (
						  children:
							- kind: R-ITEM
							  nature: _
							  name: ""
							  children:
								- kind: R-I-NAME
								  nature: _
								  name: ""
								- kind: R-I-TYPE
								  nature: '?'
								  name: ""
								- kind: R-I-VALUE
								  nature: BINDING
								  name: .
								  children:
									- kind: B-ITEM
									  nature: IDENT
									  name: bling
									- kind: B-ITEM
									  nature: IDENT
									  name: zen
		`),
	},
}

func TestRolneArgumentParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	cc.WriteToDisk = false
	cc.ShortYAML = true
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
		assert.Equal(t, test.expectedYaml, rootYaml)
	}
}
