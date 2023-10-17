package event

import "time"

type CategoryDeleted struct {
	Name    string
	Payload interface{}
}

func NewCategoryDeleted() *CategoryDeleted {
	return &CategoryDeleted{
		Name: "CategoryDeleted",
	}
}

func (e *CategoryDeleted) GetName() string {
	return e.Name
}

func (e *CategoryDeleted) GetPayload() any {
	return e.Payload
}

func (e *CategoryDeleted) SetPayload(payload any) {
	e.Payload = payload
}

func (e *CategoryDeleted) GetDateTime() time.Time {
	return time.Now()
}