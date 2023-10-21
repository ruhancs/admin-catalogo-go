package cloud

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

//url ex: https://nome-da-bucket.s3.nome-da-regiao.amazonaws.com/nome-da-key
func UploadFileToS3(filename string, file multipart.File, s3Client *s3.S3, s3Bucket string, errorUploadChan chan string, controlChan chan struct{}) {
	defer file.Close()
	_,err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key: aws.String(filename),
		Body: file,
	})
	if err != nil {
		fmt.Println(err)
		errorUploadChan <- filename
		//liberar espaco no canal
		<-controlChan
		return
	}
	//liberar espaco no canal
	<-controlChan
}