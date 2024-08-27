package ghostimg

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func useRootValue(dirPath string, rootValue bool) string {
	if !rootValue {
		dirPath = strings.TrimPrefix(dirPath, "/")
	} else {
		if !strings.HasPrefix(dirPath, "/") {
			dirPath = "/" + dirPath
		}
	}

	return dirPath
}

func createDirIfNotExists(dirPath string, rootValue bool) string {

	dp := useRootValue(dirPath, rootValue)

	if _, err := os.Stat(dp); os.IsNotExist(err) {
		os.MkdirAll(dp, os.ModePerm)
	}

	return dp
}

func saveFileMultipart(file *multipart.FileHeader, dirPath string, fileName string) error {

	src, err := file.Open()

	if err != nil {
		return ErrFileDoesntExist
	}
	defer src.Close()

	dst, err := os.Create(dirPath + "/" + fileName)
	if err != nil {
		return ErrUploadingPhoto
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return ErrInvalidFileType
	}

	return nil
}

func verifySameExtension(extension string, handlerExtension string) error {
	if extension != handlerExtension {
		return ErrNotAnImage
	}

	return nil
}

// UploadMultipartDirect uploads a file using multipart, formName should be specified in function, not in query
func UploadMultipartDirect(w http.ResponseWriter, r *http.Request, g Ghost) error {
	var parseSize int64

	if g.MaxParseSize == 0 {
		parseSize = DefaultMaxPhotoSizeMultipart
	} else {
		parseSize = g.MaxParseSize
	}

	err := r.ParseMultipartForm(parseSize)

	if err != nil {
		return ErrParseMultipartForm
	}

	for i, formName := range g.FormNames {
		file, handler, err := r.FormFile(formName)
		if err != nil {
			logIfEnabled(g, "Error getting file: %v", err)
			return ErrFileDoesntExist
		}

		v := verifyIfImage(handler)
		if !v {
			return ErrNotAnImage
		}

		extension := strings.Split(g.FileNames[i], ".")[1]
		err = verifyExtension(extension)
		if err != nil {
			logIfEnabled(g, "You must specify a valid image extension")
			return ErrNotAnImage
		}

		handlerExtension := strings.Split(handler.Filename, ".")[1]
		err = verifyExtension(handlerExtension)
		if err != nil {
			logIfEnabled(g, "You must specify a valid image extension")
			return ErrNotAnImage
		}

		err = verifySameExtension(extension, handlerExtension)
		if err != nil {
			logIfEnabled(g, "File and handler extension are not the same")
			return ErrNotSameExtension
		}

		g.DirPath = createDirIfNotExists(g.DirPath, g.AtRootOfDirectory)

		err = saveFileMultipart(handler, g.DirPath, g.FileNames[i])
		if err != nil {
			logIfEnabled(g, "Error saving file: %v", err)
			return ErrUploadingPhoto
		}

		logIfEnabled(g, "Uploaded file: %s as %s in %s", handler.Filename, g.FileNames[i], g.DirPath)
		defer file.Close()
	}

	return nil
}

func verifyIfImage(file *multipart.FileHeader) bool {
	return strings.Contains(file.Header.Get("Content-Type"), "image")
}

func verifyExtension(extension string) error {
	if extension != "jpg" && extension != "jpeg" && extension != "png" && extension != "gif" {
		return ErrNotAnImage
	}

	return nil
}

// ServeImage serves an image from a directory
func ServeImage(w http.ResponseWriter, r *http.Request, i Img) error {

	extension := strings.Split(i.FileName, ".")[1]
	err := verifyExtension(extension)

	if err != nil {
		logIfEnabled(i, "You must specify a valid image extension")
		return ErrNotAnImage
	}

	i.DirPath = useRootValue(i.DirPath, i.AtRootOfDirectory)

	if _, err := os.Stat(i.DirPath); os.IsNotExist(err) {
		logIfEnabled(i, "Directory does not exist")
		return ErrDirDoesntExist
	}

	file, err := os.Open(i.DirPath + "/" + i.FileName)

	if err != nil {
		logIfEnabled(i, "Error opening file: %v", err)
		return ErrFileDoesntExist
	}

	defer file.Close()

	extension = getExtension(extension)
	w.Header().Set("Content-Type", "image/"+extension)

	_, err = io.Copy(w, file)

	if err != nil {
		logIfEnabled(i, "Error copying file: %v", err)
		return ErrUploadingPhoto
	}

	return nil
}

func getExtension(extension string) string {
	if extension == "jpg" {
		return "jpeg"
	} else {
		return extension
	}
}

// Uploads a binary file from a client
func UploadBinary(w http.ResponseWriter, r *http.Request, g Ghost) error {

	if g.FileNames[0] == "" {
		logIfEnabled(g, "You must specify a file name")
		return ErrFileDoesntExist
	}

	extensionFileName := strings.Split(g.FileNames[0], ".")[1]

	err := verifyExtension(extensionFileName)
	if err != nil {
		logIfEnabled(g, "You must specify a valid image extension")
		return ErrNotAnImage
	}

	var file io.ReadCloser = r.Body

	bodyFile, err := io.ReadAll(file)

	defer r.Body.Close()

	if err != nil {
		logIfEnabled(g, "Error reading file: %v", err)
		return ErrInvalidFileType
	}

	verifyPhotoExtension := verifyIfImageBinary(bodyFile)
	if !verifyPhotoExtension {
		logIfEnabled(g, "File is not an image")
		return ErrNotAnImage
	}

	dirpath := createDirIfNotExists(g.DirPath, g.AtRootOfDirectory)
	dst, err := os.Create(dirpath + "/" + g.FileNames[0])
	if err != nil {
		logIfEnabled(g, "Error creating file: %v", err)
		return ErrUploadingPhoto
	}

	defer dst.Close()

	_, err = dst.Write(bodyFile)
	if err != nil {
		logIfEnabled(g, "Error writing file: %v", err)
		return ErrUploadingPhoto
	}

	logIfEnabled(g, "Uploaded file: %s as %s in %s", g.FileNames[0], g.FileNames[0], dirpath)

	return nil
}

func verifyIfImageBinary(data []byte) bool {
	contentType := http.DetectContentType(data)
	return strings.HasPrefix(contentType, "image")
}

func logIfEnabled(g Loggable, message string, args ...interface{}) {
	if !g.DisableLogging() {
		log.Printf(message, args...)
	}
}
