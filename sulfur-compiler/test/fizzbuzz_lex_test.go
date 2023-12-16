package test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sulfur-compiler/compiler"
	"sulfur-compiler/context"
	"testing"
)

func TestFizzBuzzLexingMain(t *testing.T) {
	//
	// arrange
	//
	target := "fizzbuzz"
	cc := context.NewCompilerContext(target)
	cc.StopPoint = "lex_target"
	targetYamlFileLocation := cc.StagingDir + "/" + context.LEXING_PREFIX_EN + "/file-0000.token.yaml"
	expectedYaml, err := os.ReadFile("fizzbuzz_lex_test.main.yaml")
	expectedYamlStr := string(expectedYaml)

	//
	// act
	//
	err = compiler.Compiler(cc)

	//
	// assert
	//
	if err != nil {
		t.Errorf("lexing main failed, error: %v", err)
	}
	result, err := os.ReadFile(targetYamlFileLocation)
	resultStr := string(result)
	assert.Equal(t, resultStr, expectedYamlStr, "tokenization does not match")
}
