package ghostimg

type Ghost struct {
	DirPath           string   // directory path where the file will be saved
	FormNames         []string // formName specified in client
	FileNames         []string // name of files saved
	AtRootOfDirectory bool     // if true, the file will be saved at the root of the directory, if false, it will be saved at root of project
	MaxParseSize      int64    // maximum size of file that can be parsed
	DisableLog        bool     // if true, the log will be disabled
}

type Img struct {
	DirPath           string // directory path where the file is saved
	FileName          string // name of file to be retrieved
	AtRootOfDirectory bool   // if true, the file will be saved at the root of the directory, if false, it will be saved at root of project
	DisableLog        bool   // if true, the log will be disabled
}

func (g Ghost) DisableLogging() bool {
	return g.DisableLog
}

func (i Img) DisableLogging() bool {
	return i.DisableLog
}

type Loggable interface {
	DisableLogging() bool
}
