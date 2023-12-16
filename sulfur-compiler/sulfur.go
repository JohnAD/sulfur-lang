// The sulfur transpiler

package main

import (
	"flag"
	"log"
	"os"
	"sulfur-compiler/compiler"
	"sulfur-compiler/context"
)

func main() {

	rootDirPtr := flag.String("root", ".", "the root directory to begin compilation in")
	stopCompilationStagePtr := flag.String("stop", "", "the stage in which to stop compilation; defaults to full compilation")
	flag.Parse()

	target := "main"
	if len(flag.Args()) != 0 {
		target = flag.Args()[0]
	}
	cc := context.NewCompilerContext(target)
	cc.RootDir = *rootDirPtr
	cc.StopPoint = *stopCompilationStagePtr

	err := compiler.Compiler(cc)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

//fmt.Println("Tokens:")
//for _, t := range tokens {
//fmt.Printf("%v ", t)
//}
//fmt.Println("\n\nRebuild:\n")
//rebuild := lexer.RebuildFromTokens(tokens)
//fmt.Println(rebuild)

//fmt.Println("AST: ")
//rootYamlBytes, _ := yaml.Marshal(ast)
//fmt.Println(string(rootYamlBytes))
