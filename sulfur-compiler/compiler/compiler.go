package compiler

import (
	"sulfur-compiler/context"
	"sulfur-compiler/lexer"
	"sulfur-compiler/sparser"
)

func Compiler(cc context.CompilerContext) error {
	//
	// 0: staging
	//
	err := context.CleanStagingDir(&cc)
	if err != nil {
		return err
	}
	err = context.CreateLexParseRoundDir(&cc)
	if err != nil {
		return err
	}
	if cc.StopPoint == "stage" {
		return err
	}

	//
	// 1: lex target
	//
	err, target_tokens := lexer.LexFile(&cc, cc.Target)
	if err != nil {
		return err
	}
	if cc.StopPoint == "lex_target" {
		return nil
	}

	//
	// 2: parse target
	//
	err, _ = sparser.ParseTokensToAst(&cc, target_tokens)
	if err != nil {
		return err
	}
	if cc.StopPoint == "parse_target" {
		return nil
	}

	//
	// 3: recursively lex/parse sub-targets
	//
	return nil
}
