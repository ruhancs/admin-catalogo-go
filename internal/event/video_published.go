package event

import "time"

type VideoPublish struct {
	Name    string
	Payload interface{}
}

func NewVideoPublish() *VideoPublish {
	return &VideoPublish{
		Name: "VideoPublished",
	}
}

func (e *VideoPublish) GetName() string {
	return e.Name
}

func (e *VideoPublish) GetPayload() any {
	return e.Payload
}

func (e *VideoPublish) SetPayload(payload any) {
	e.Payload = payload
}

func (e *VideoPublish) GetDateTime() time.Time {
	return time.Now()
}