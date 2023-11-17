package cloud

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteFileInBucket(filename,bucketName  string, s3Client *s3.S3,) error{
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(filename),
	}
	_,err := s3Client.DeleteObject(input)
	if err != nil {
		return err
	}
	
	return nil
}