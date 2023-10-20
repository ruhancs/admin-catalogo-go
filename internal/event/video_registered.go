package event

import "time"

type VideoRegistered struct {
	Name    string
	Payload interface{}
}

func NewVideoRegistered() *VideoRegistered {
	return &VideoRegistered{
		Name: "VideoRegistered",
	}
}

func (e *VideoRegistered) GetName() string {
	return e.Name
}

func (e *VideoRegistered) GetPayload() any {
	return e.Payload
}

func (e *VideoRegistered) SetPayload(payload any) {
	e.Payload = payload
}

func (e *VideoRegistered) GetDateTime() time.Time {
	return time.Now()
}