package ghostimg

import (
	"mime/multipart"
	"net/textproto"
	"testing"
)

func TestPath(t *testing.T) {

	pathmodel := Ghost{
		DirPath:           "app/utils",
		FileNames:         []string{"test.png"},
		FormNames:         []string{"test", "fab"},
		AtRootOfDirectory: true,
	}

	if pathmodel.DirPath != "app/utils" {
		t.Errorf("Expected %s, got %s", "app/utils", pathmodel.DirPath)
	}

}

func TestMultipartRequest(t *testing.T) {
	// var b bytes.Buffer

	// w := multipart.NewWriter(&b)

	// req := httptest.NewRequest(http.MethodPost, "/upload?formName=test", &b)

	// req.Header.Set("Content-Type", w.FormDataContentType())

	// // rr := httptest.NewRecorder()

	// w.Close()

	g := Ghost{
		DirPath:           "app/utils",
		FileNames:         []string{"test.png", "salt.png"},
		FormNames:         []string{"test", "fab"},
		AtRootOfDirectory: true,
	}
	file := &multipart.FileHeader{
		Filename: "test.png",
		Header: textproto.MIMEHeader{
			"Content-Type": []string{"image/png"},
		},
		Size: 100,
	}

	err := saveFileMultipart(file, g.DirPath, g.FileNames[0])

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
