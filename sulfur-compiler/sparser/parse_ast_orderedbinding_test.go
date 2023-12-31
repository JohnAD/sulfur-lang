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
			kind: ROOT
			nature: _
			name: ""
			children:
				- kind: STATEMENT
				  nature: '?'
				  name: ""
				  children:
					- kind: S-ITEM
					  nature: BINDING
					  name: .
					  children:
						- kind: B-ITEM
						  nature: IDENT
						  name: foo
						- kind: B-ITEM
						  nature: IDENT
						  name: bar
					- kind: S-ITEM
					  nature: IDENT
					  name: next
			`),
	}, {
		"deeper bind", // TODO: next
		helpers.Dedent(`
			foo.bar\a.b.c\x.z
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
					  nature: BINDING
					  name: .
					  children:
						- kind: B-ITEM
						  nature: IDENT
						  name: foo
						- kind: B-ITEM
						  nature: IDENT
						  name: bar
					- kind: S-ITEM
					  nature: IDENT
					  name: next
			`),
	},
}

func TestOrderedBindingParsing(t *testing.T) {
	cc := context.NewCompilerContext("main")
	cc.WriteToDisk = false
	cc.ShortYAML = true
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
