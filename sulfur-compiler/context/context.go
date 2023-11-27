package context

type FileDetail struct {
	fileId   int
	filepath string
}

type CompilerContext struct {
	RootDir       string
	Target        string
	StagingDir    string
	Files         []FileDetail
	SaveLexedFlag bool
}

func NewCompilerContext(mainTarget string) CompilerContext {
	cc := CompilerContext{
		RootDir:       ".",
		Target:        mainTarget,
		StagingDir:    "_build",
		Files:         []FileDetail{},
		SaveLexedFlag: true,
	}
	return cc
}

func SetNewFile(cc *CompilerContext, target string) int {
	fileId := len(cc.Files)
	newFile := FileDetail{fileId, target}
	cc.Files = append(cc.Files, newFile)
	return fileId
}
