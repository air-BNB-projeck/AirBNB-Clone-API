package helper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploaderS3() *s3.S3 {
	s3Config := &aws.Config{
		Region: aws.String("ap-southeast-3"),
		Credentials: credentials.NewStaticCredentials("AKIAVU6WN7VLMK6VYS6A", "zxQwQr5PhaF0nDAhOdukYa7zaw8O5s4ahBgqPsCe", ""),
	}
	s3Session, _ := session.NewSession(s3Config)

	uploader := s3.New(s3Session)

	return uploader
}