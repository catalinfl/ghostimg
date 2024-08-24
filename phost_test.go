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

	err := SaveFileMultipart(&file, Paths{
		DirPath:  "app/utils",
		FilePath: "test.png",
	})

	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

}

func TestMultipartRequest(t *testing.T) {
	// var b bytes.Buffer

	// w := multipart.NewWriter(&b)

	// req := httptest.NewRequest(http.MethodPost, "/upload?formName=test", &b)

	// req.Header.Set("Content-Type", w.FormDataContentType())

	// // rr := httptest.NewRecorder()

	// w.Close()

	p := Paths{
		DirPath:  "app/utils",
		FilePath: "test.png",
	}

	file := &multipart.FileHeader{
		Filename: "test.png",
		Header: textproto.MIMEHeader{
			"Content-Type": []string{"image/png"},
		},
		Size: 100,
	}

	err := SaveFileMultipart(file, p)

	if err != nil {
		t.Errorf("%s", err)
	}

	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("Expected %d, got %d", http.StatusOK, status)
	// }

	// expected := "File uploaded successfully\n"

	// if rr.Body.String() != expected {
	// 	t.Errorf("Expected %s, got %s", expected, rr.Body.String())
	// }

}
