package ghostimg

import (
	"mime/multipart"
	"net/http"
)

func SaveMultipartForm(file *multipart.FileHeader) error {
	headerType := file.Header.Get("Content-Type")

	if headerType != "image/png" && headerType != "image/jpeg" {
		return ErrInvalidFileType
	}

	return nil
}

func SaveMultipart(formName string, w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(DefaultMaxPhotoSizeMultipart)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile(formName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	err = SaveMultipartForm(handler)
}

func UploadFileMultipart(w http.ResponseWriter, r *http.Request) {
	formName := r.URL.Query().Get("formName")

	if formName == "" {
		http.Error(w, "formName is required", http.StatusBadRequest)
		return
	}

	SaveMultipart(formName, w, r)
}
