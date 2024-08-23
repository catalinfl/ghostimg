package ghostimg

import (
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/catalinfl/ghostimg/utils"
)

func TestPath(t *testing.T) {

	pathmodel := Paths{
		DirPath:  "app/utils",
		FilePath: "test.png",
	}

	path := utils.GetRealPath(pathmodel.DirPath, pathmodel.FilePath)

	t.Logf(path)

	if path != "/app/utils/test.png" {
		t.Errorf("Expected /app/utils/test.png, got %s", path)
	}

}

func TestMultipart(t *testing.T) {

	file := multipart.FileHeader{
		Filename: "test.png",
		Header: textproto.MIMEHeader{
			"Content-Type": []string{"image/png"},
		},
		Size: 100,
	}

	err := SaveMultipartForm(&file)

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

}
