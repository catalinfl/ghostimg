package ghostimg

import "errors"

var (
	ErrInvalidFileType      = errors.New("invalid file type")
	ErrFormNameNotSpecified = errors.New("formname is required in url, use ?formname=")
	ErrParseMultipartForm   = errors.New("error parsing multipart form")
	ErrFileDoesntExist      = errors.New("file does not exist in form")
	ErrUploadingPhoto       = errors.New("error uploading photo")
	ErrNotAnImage           = errors.New("file is not an image")
	ErrNotSameExtension     = errors.New("file and handler extension are not the same")
	ErrDirDoesntExist       = errors.New("directory does not exist")
)
