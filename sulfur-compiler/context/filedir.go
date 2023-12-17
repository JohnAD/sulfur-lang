package context

import (
	"os"
	"strings"
)

const LEXING_PREFIX_EN = "00_Lexing"
const PARSING_PREFIX_EN = "01_Parsing"

func CleanStagingDir(cc *CompilerContext) error {
	stagingPath := buildStagePath(cc, "")
	err := deleteDirectoryIfExists(stagingPath)
	if err != nil {
		return err
	}
	err = os.Mkdir(stagingPath, 0750)
	return err
}

func buildStagePath(cc *CompilerContext, stage string) string {
	fullBuildDir := cc.RootDir + "/" + cc.StagingDir
	if stage == "" {
		return fullBuildDir
	}
	finalBuildDir := fullBuildDir + "/" + stage
	return finalBuildDir
}

// snagged from https://www.tutorialspoint.com/golang-program-to-delete-empty-and-non-empty-directory
func deleteDirectoryIfExists(path string) error {
	dr, err := os.Open(path) //open the path using os package function
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") == true {
			return nil
		}
		return err
	}
	defer dr.Close()
	names, err := dr.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(path + "/" + name) //remove the path
		if err != nil {
			return err
		}
	}
	return os.Remove(path)
}

//
// LEXER RESULTS
//

func CreateLexResultDir(cc *CompilerContext) error {
	path := GetLexResultPath(cc)
	err := os.Mkdir(path, 0750)
	return err
}
func GetLexResultPath(cc *CompilerContext) string {
	lexRoundPath := buildStagePath(cc, LEXING_PREFIX_EN)
	return lexRoundPath
}

//
// PARSER RESULTS
//

func CreateParseResultDir(cc *CompilerContext) error {
	err := os.Mkdir(GetParseResultPath(cc), 0750)
	return err
}
func GetParseResultPath(cc *CompilerContext) string {
	parseRoundPath := buildStagePath(cc, PARSING_PREFIX_EN)
	return parseRoundPath
}
