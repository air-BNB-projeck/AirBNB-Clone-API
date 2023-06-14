package helper

import (
	"alta/air-bnb/app/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploaderS3() *s3.S3 {
	appConfig := config.ReadEnv()
	s3Config := &aws.Config{
		Region: aws.String("ap-southeast-3"),
		Credentials: credentials.NewStaticCredentials(appConfig.AWS_ACCESS_KEY, appConfig.AWS_SECRET_KEY, ""),
	}
	s3Session, _ := session.NewSession(s3Config)

	uploader := s3.New(s3Session)

	return uploader
}