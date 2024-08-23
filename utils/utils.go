package utils

import (
	"strings"
)

func GetRealPath(dirPath string, filePath string) string {

	path := ""

	if strings.HasPrefix(dirPath, "/") {
		path = dirPath
	} else {
		path = "/" + dirPath
	}

	path = strings.TrimSuffix(path, "/")

	if strings.HasPrefix(filePath, "/") {
		path = path + filePath
	} else {
		path = path + "/" + filePath
	}

	return path

}
