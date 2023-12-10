package sparser

import (
	"gopkg.in/yaml.v3"
	"sulfur-compiler/context"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
	"testing"
)

var parseRolneParameterTests = []parseTest{
	{
		"simple parameter",
		helpers.Dedent(`
			parameters { bling::String }
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
					  name: parameters
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
							  nature: IDENT
							  name: bling
							  children: []
							- type: R-I-TYPE
							  nature: IDENT
							  name: String
							  children: []
							- type: R-I-VALUE
							  nature: '?'
							  name: ""
							  children: []
		`),
	}, {
		"parameter with default",
		helpers.Dedent(`
			parameters { bling::String = "joe" }
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
					  name: parameters
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
							  nature: IDENT
							  name: bling
							  children: []
							- type: R-I-TYPE
							  nature: IDENT
							  name: String
							  children: []
							- type: R-I-VALUE
							  nature: STR-LIT
							  name: joe
							  children: []
		`),
	},
}

func TestRolneParameterParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	for _, test := range parseRolneParameterTests {
		err, tokens := lexer.LexString(test.source)
		if err != nil {
			t.Errorf("before parse test even started, failed to lex on test `%s` LEX:\n%v\nERR: %s", test.desc, tokens, err.Error())
			return
		}
		err, ast := ParseTokensToAst(&cc, tokens)
		rootYamlBytes, _ := yaml.Marshal(ast)
		rootYaml := string(rootYamlBytes)
		if err != nil {
			t.Errorf(">>>> on test `%s`, parser error: %v", test.desc, err)
		}
		if rootYaml != test.expectedYaml {
			t.Errorf(">>>> on test `%s`\nPARSING>>>\n%s<<<EOF\n\nGOT>>>\n%s<<<EOF\n\nEXPECTED>>>\n%s<<<EOF", test.desc, test.source, rootYaml, test.expectedYaml)
		}
	}
}
