package event

import "time"

type CategoryCreated struct {
	Name    string
	Payload interface{}
}

func NewCategoryCreated() *CategoryCreated {
	return &CategoryCreated{
		Name: "CategoryCreated",
	}
}

func (e *CategoryCreated) GetName() string {
	return e.Name
}

func (e *CategoryCreated) GetPayload() any {
	return e.Payload
}

func (e *CategoryCreated) SetPayload(payload any) {
	e.Payload = payload
}

func (e *CategoryCreated) GetDateTime() time.Time {
	return time.Now()
}