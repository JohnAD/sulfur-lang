package context

type FileDetail struct {
	fileId   int
	filepath string
}

type CompilerContext struct {
	RootDir     string
	Target      string
	StagingDir  string
	Files       []FileDetail
	StopPoint   string
	WriteToDisk bool
	ShortYAML   bool
}

func NewCompilerContext(mainTarget string) CompilerContext {
	cc := CompilerContext{
		RootDir:     ".",
		Target:      mainTarget,
		StagingDir:  "_build",
		Files:       []FileDetail{},
		StopPoint:   "",
		WriteToDisk: true, // set to false by unit tests
		ShortYAML:   false,
	}
	return cc
}

func SetNewFile(cc *CompilerContext, target string) int {
	fileId := len(cc.Files)
	newFile := FileDetail{fileId, target}
	cc.Files = append(cc.Files, newFile)
	return fileId
}
