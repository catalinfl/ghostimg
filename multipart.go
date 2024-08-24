package ghostimg

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func SaveFileMultipart(file *multipart.FileHeader, p Paths) error {

	var err error

	if os.Stat(p.DirPath); os.IsNotExist(err) {
		err := os.MkdirAll(p.DirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(p.FilePath)
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
func UploadMultipartDirect(w http.ResponseWriter, r *http.Request, p Paths, formNames ...string) error {
	err := r.ParseMultipartForm(DefaultMaxPhotoSizeMultipart)

	if err != nil {
		return ErrParseMultipartForm
	}

	for _, formName := range formNames {
		file, handler, err := r.FormFile(formName)

		log.Printf("formName: %s", formName)

		if err != nil {
			return ErrFileDoesntExist
		}

		err = SaveFileMultipart(handler, p)

		if err != nil {
			return ErrUploadingPhoto
		}

		log.Printf("Uploaded file: %s", handler.Filename)

		defer file.Close()
	}

	return nil
}

// UploadFileMultipart uploads a file using multipart form data. The form name must be specified in the query string. (?formName=...)
func UploadFileMultipart(w http.ResponseWriter, r *http.Request, p Paths) error {
	formName := r.URL.Query().Get("formName")

	log.Printf("formName: %s", formName)

	if formName == "" {
		return ErrFormNameNotSpecified
	}

	UploadMultipartDirect(w, r, p, formName)

	return nil
}
