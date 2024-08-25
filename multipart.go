package ghostimg

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func createDirIfNotExists(dirPath string, rootValue bool) string {

	if !rootValue {
		dirPath = strings.TrimPrefix(dirPath, "/")
	} else {
		if !strings.HasPrefix(dirPath, "/") {
			dirPath = "/" + dirPath
		}
	}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Printf("Error creating directory: %v", err)
		}
	}

	return dirPath
}

func saveFileMultipart(file *multipart.FileHeader, dirPath string, fileName string) error {

	src, err := file.Open()

	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dirPath + "/" + fileName)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return err
	}

	return nil
}

// UploadMultipartDirect uploads a file using multipart, formName should be specified in function, not in query
func UploadMultipartDirect(w http.ResponseWriter, r *http.Request, g Ghost) error {
	err := r.ParseMultipartForm(DefaultMaxPhotoSizeMultipart)

	if err != nil {
		return ErrParseMultipartForm
	}

	for i, formName := range g.FormNames {
		file, handler, err := r.FormFile(formName)
		if err != nil {
			log.Printf("Error retrieving form file for form name %s: %v", formName, err)
			return ErrFileDoesntExist
		}

		g.DirPath = createDirIfNotExists(g.DirPath, g.AtRootOfDirectory)

		err = saveFileMultipart(handler, g.DirPath, g.FileNames[i])
		if err != nil {
			log.Printf("Error saving file: %v", err)
			return ErrUploadingPhoto
		}

		log.Printf("Uploaded file: %s as %s in %s", handler.Filename, g.FileNames[i], g.DirPath)
		defer file.Close()
	}

	return nil
}
