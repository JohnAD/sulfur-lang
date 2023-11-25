// The sulfur transpiler

package main

import (
	"fmt"
	"log"
	"sulfur-compiler/lexer"
)

func main() {
	base_dir := "/home/johnd/Projects/sulfur-lang/advent-src"
	target := "main.sulfur"
	err, tokens := lexer.LexFile(base_dir, target)
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
}
