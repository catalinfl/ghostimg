package ghostimg

type Ghost struct {
	DirPath           string   // directory path where the file will be saved
	FormNames         []string // formName specified in client
	FileNames         []string // name of files saved
	AtRootOfDirectory bool     // if true, the file will be saved at the root of the directory, if false, it will be saved at root of project
}
