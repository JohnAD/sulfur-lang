// The sulfur transpiler

package main

import (
	"fmt"
	"log"
	"sulfur-compiler/context"
	"sulfur-compiler/lexer"
	"sulfur-compiler/parser"
)

func main() {
	cc := context.NewCompilerContext("main")
	cc.RootDir = "/home/johnd/Projects/sulfur-lang/advent-src"
	target := "main"
	err := context.CleanStagingDir(&cc)
	if err != nil {
		log.Fatal(err)
	}

	// round 1: Dynamic Lexing/Parsing
	err = context.CreateLexParseRoundDir(&cc)
	if err != nil {
		log.Fatal(err)
	}
	err, tokens := lexer.LexFile(&cc, target)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Tokens:")
	for _, t := range tokens {
		fmt.Printf("%v ", t)
	}
	fmt.Println("\n\nRebuild:\n")
	rebuild := lexer.RebuildFromTokens(tokens)
	fmt.Println(rebuild)

	err, ast := parser.ParseTokensToAst(&cc, tokens)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AST:")
	fmt.Printf("%v ", ast)
}
