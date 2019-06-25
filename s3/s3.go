package s3

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/yanshiyason/speren/conf"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Service wraps a s3 client and a bucket name
type Service struct {
	client   *s3.S3
	Bucket   string
	uploader *s3manager.Uploader
}

// NewServiceFromConf initializes an instance from configuration file.
func NewServiceFromConf() *Service {
	config, err := conf.NewService()
	if err != nil {
		panic(err)
	}
	bucket := config.String("s3.bucket")
	region := config.String("aws.region")
	accessKey := config.String("aws.access_key_id")
	secret := config.String("aws.secret_access_key")

	service, err := NewService(region, bucket, accessKey, secret)
	if err != nil {
		panic(err)
	}

	return service
}

// NewService initialize a s3 client
func NewService(region, bucket, accessKeyID, secret string) (*Service, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      &region,
		Credentials: credentials.NewStaticCredentials(accessKeyID, secret, ""),
	})
	if err != nil {
		return nil, err
	}
	client := s3.New(sess)
	uploader := s3manager.NewUploader(sess)

	return &Service{
		client:   client,
		Bucket:   bucket,
		uploader: uploader,
	}, nil
}

// Store stores a base64 encoded blob into s3 and returns it's url
// implements the FileStorage interface
func (s *Service) Store(filename string, blob []byte) (string, error) {
	// str, err := base64.StdEncoding.EncodeToString(blob)
	// bb, err := base64.StdEncoding.DecodeString(base64Str)
	// if err != nil {
	// 	return nil, err
	// }
	// return s.upload(filename, bytes.NewReader(bb))
	return s.upload(filename, bytes.NewReader(blob))
}

// uploads the file to s3 and returns the url
func (s *Service) upload(path string, body io.Reader) (string, error) {
	mimeType, err := getFileContentType(body)
	if err != nil {
		return "", err
	}

	output, err := s.uploader.Upload(&s3manager.UploadInput{
		Bucket:      &s.Bucket,
		Key:         &path,
		ContentType: &mimeType,
		Body:        body,
	})
	if err != nil {
		return "", err
	}

	location, err := url.Parse(output.Location)
	if err != nil {
		return "", err
	}

	return location.String(), nil
}

// getFileContentType figures out the content type
// https://golangcode.com/get-the-content-type-of-file/
func getFileContentType(out io.Reader) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
