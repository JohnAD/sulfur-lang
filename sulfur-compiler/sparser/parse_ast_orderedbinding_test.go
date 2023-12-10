package sparser

import (
	"gopkg.in/yaml.v3"
	"sulfur-compiler/context"
	"sulfur-compiler/helpers"
	"sulfur-compiler/lexer"
	"testing"
)

var parseOrderedBindingTests = []parseTest{
	{
		"simple bind",
		helpers.Dedent(`
			foo.bar next
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
					  nature: BINDING
					  name: .
					  children:
						- type: BINDING-CHILD
						  nature: IDENT
						  name: foo
						  children: []
						- type: BINDING-CHILD
						  nature: IDENT
						  name: bar
						  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: next
					  children: []
			`),
	}, {
		"deeper bind", // TODO: next
		helpers.Dedent(`
			foo.bar.a.b
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
					  nature: BINDING
					  name: .
					  children:
						- type: BINDING-CHILD
						  nature: IDENT
						  name: foo
						  children: []
						- type: BINDING-CHILD
						  nature: IDENT
						  name: bar
						  children: []
					- type: S-ITEM
					  nature: IDENT
					  name: next
					  children: []
			`),
	},
}

func TestOrderedBindingParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	for _, test := range parseOrderedBindingTests {
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
