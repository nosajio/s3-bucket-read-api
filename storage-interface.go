package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

// CreateAWSSession initiates a connection through the AWS SDK, then returns it
func CreateAWSSession(awsRegion string) *s3.S3 {
	s := s3.New(session.New(), &aws.Config{
		Region: aws.String(awsRegion),
	})
	return s
}

// ListFilesForBucket returns a list of urls for all files with an extension
// included in `filetypes`
func ListFilesForBucket(
	s *s3.S3,
	bucketName string,
	filetypes []string) *s3.ListObjectsOutput {
	input := &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	}
	result, err := s.ListObjects(input)
	if err != nil {
		Error.Println(err)
		return nil
	}
	return result
}

// GetObjectBytes downloads the specified file and returns the bytes
func GetObjectBytes(
	s *s3.S3,
	bucketName string,
	filename string) (string, []byte) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	}
	result, err := s.GetObject(input)
	if err != nil {
		Error.Println(fmt.Sprintf("filekey: %s", *aws.String(filename)))
		Error.Println(err)
		return "", nil
	}
	defer result.Body.Close()
	buffer := bytes.NewBuffer(nil)
	contentType := *result.ContentType
	if _, err := io.Copy(buffer, result.Body); err != nil {
		Error.Println(err)
		return contentType, nil
	}

	return contentType, buffer.Bytes()
}
