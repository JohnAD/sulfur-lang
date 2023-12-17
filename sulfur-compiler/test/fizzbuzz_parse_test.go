package test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"sulfur-compiler/compiler"
	"sulfur-compiler/context"
	"testing"
)

func TestFizzBuzzParsingMain(t *testing.T) {
	//
	// arrange
	//
	target := "fizzbuzz"
	cc := context.NewCompilerContext(target)
	cc.StopPoint = "parse_target"
	targetYamlFileLocation := cc.StagingDir + "/" + context.LEXING_PREFIX_EN + "/file-0000.token.yaml"
	expectedYaml, err := os.ReadFile("fizzbuzz_parse_test.main.yaml")
	expectedYamlStr := string(expectedYaml)
	//
	// act
	//
	err = compiler.Compiler(cc)
	//
	// assert
	//
	if err != nil {
		t.Errorf("%s failed, error: %v", cc.StopPoint, err)
		return
	}
	result, err := os.ReadFile(targetYamlFileLocation)
	resultStr := string(result)
	assert.Equal(t, resultStr, expectedYamlStr)
}
