package ghostimg

import (
	"log"
	"mime/multipart"
	"net/http"
)

func SaveFileMultipart(file *multipart.FileHeader) error {
	return nil
}

// UploadMultipartDirect uploads a file using multipart, formName should be specified in function, not in query
func UploadMultipartDirect(w http.ResponseWriter, r *http.Request, formNames ...string) error {

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

		err = SaveFileMultipart(handler)

		if err != nil {
			return ErrUploadingPhoto
		}

		log.Printf("Uploaded file: %s", handler.Filename)

		defer file.Close()
	}

	return nil
}

// UploadFileMultipart uploads a file using multipart form data. The form name must be specified in the query string. (?formName=...)
func UploadFileMultipart(w http.ResponseWriter, r *http.Request) error {

	formName := r.URL.Query().Get("formName")

	log.Printf("formName: %s", formName)

	if formName == "" {
		return ErrFormNameNotSpecified
	}

	UploadMultipartDirect(w, r, formName)

	return nil
}
