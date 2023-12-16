package context

import (
	"os"
	"strings"
)

const LEXING_PREFIX_EN = "00_Lexing"

func CleanStagingDir(cc *CompilerContext) error {
	stagingPath := buildStagePath(cc, "")
	err := deleteDirectoryIfExists(stagingPath)
	if err != nil {
		return err
	}
	err = os.Mkdir(stagingPath, 0750)
	return err
}

func CreateLexParseRoundDir(cc *CompilerContext) error {
	err := os.Mkdir(GetLexParseRoundPath(cc), 0750)
	return err
}
func GetLexParseRoundPath(cc *CompilerContext) string {
	lexParseRoundPath := buildStagePath(cc, LEXING_PREFIX_EN)
	return lexParseRoundPath
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
