package filesystem_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/yanshiyason/speren/filesystem"
)

func dummyImg(filename string) []byte {
	path := filepath.Join("testdata", filename)
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return blob
}

func tearDown() {
	os.RemoveAll("../public/uploads")
	os.MkdirAll("../public/uploads", os.ModePerm)
}

func Test_Filesystem_Store(t *testing.T) {
	defer tearDown()
	storage := &filesystem.Storage{}

	cases := []struct {
		in, out string
	}{
		{"dummy.jpg", "/public/uploads/dummy.jpg"},
		{"dummy.gif", "/public/uploads/dummy.gif"},
		{"dummy.png", "/public/uploads/dummy.png"},
	}
	for _, c := range cases {
		url, err := storage.Store(c.in, dummyImg(c.in))

		// happy path (unexpected error)
		if err != nil {
			t.Fatalf(`Unexpected error: %s`, err)
		}

		if url != c.out {
			t.Errorf(`
			Expected URL: %s
			Actual URL:   %s
			`, c.out, url)
		}
	}
}
