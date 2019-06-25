package s3_test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/yanshiyason/speren/s3"

	"gopkg.in/h2non/gock.v1"
)

var service *s3.Service

func init() {
	bucket := "test_bucket"
	region := "ap-northeast-1"
	accessKey := "test_accessKey"
	secret := "test_secret"

	var err error
	service, err = s3.NewService(region, bucket, accessKey, secret)
	if err != nil {
		panic(err)
	}
}

func dummyImg(filename string) []byte {
	path := filepath.Join("testdata", filename)
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return blob
}

type RequestFailure = awserr.RequestFailure

func Test_S3_Store(t *testing.T) {
	cases := []struct {
		code     int
		url      string
		filename string
		mimeType string
		err      string
	}{
		{200, "https://s3.ap-northeast-1.amazonaws.com/test_bucket/dummy.jpg", "dummy.jpg", "^image/jpeg$", ""},
		{200, "https://s3.ap-northeast-1.amazonaws.com/test_bucket/dummy.gif", "dummy.gif", "^image/gif$", ""},
		{200, "https://s3.ap-northeast-1.amazonaws.com/test_bucket/dummy.png", "dummy.png", "^image/png$", ""},
		{400, "", "dummy.jpg", "image/jpeg", "BadRequest"},
	}
	for _, c := range cases {
		gock.New("https://s3.ap-northeast-1.amazonaws.com/test_bucket").
			Put("/"+c.filename).
			MatchHeader("Content-Type", c.mimeType).
			Reply(c.code)

		url, err := service.Store(c.filename, dummyImg(c.filename))

		// happy path (unexpected error)
		if c.err == "" && err != nil {
			if c.err != err.Error() {
				t.Errorf(`
			Expected no error: %s
			Actual error:      %s
			`, c.err, err)
			}
		}

		// sad path (expected error)
		if c.err != "" {
			if !strings.Contains(fmt.Sprintf("%s", err), c.err) {
				t.Errorf(`
			Expected error to contain: %s
			Actual error was:          %s
			`, c.err, err)
			}
		}

		if url != c.url {
			t.Errorf(`
			Expected URL: %s
			Actual URL:   %s
			`, c.url, url)
		}
	}
}
