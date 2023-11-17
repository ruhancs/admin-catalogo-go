package event

import "time"

type VideoFileUploaded struct {
	Name    string
	Payload interface{}
}

func NewVideoFileUploaded() *VideoFileUploaded {
	return &VideoFileUploaded{
		Name: "VideoFileUploaded",
	}
}

func (e *VideoFileUploaded) GetName() string {
	return e.Name
}

func (e *VideoFileUploaded) GetPayload() any {
	return e.Payload
}

func (e *VideoFileUploaded) SetPayload(payload any) {
	e.Payload = payload
}

func (e *VideoFileUploaded) GetDateTime() time.Time {
	return time.Now()
}
