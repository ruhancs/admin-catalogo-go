package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	ID string `json:"id" valid:"required"`
	Name string `json:"name" valid:"required,stringlength(4|25)"`
	Description string `json:"description" valid:"required"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCategory(name,description string) (*Category,error) {
	id := uuid.NewV4().String()
	createdAt := time.Now()
	isActive := true
	category := &Category{
		ID: id,
		Name: name,
		Description: description,
		IsActive: isActive,
		CreatedAt: createdAt,
	}

	err := category.Validate()
	if err != nil {
		return nil,err
	}

	return category,nil
}

func (c *Category) Active() {
	c.IsActive = true
}

func (c *Category) DeActive() {
	c.IsActive = false
}

func(c *Category) Update(name, description string) error {
	c.Name = name
	c.Description = description
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (c *Category) Validate() error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	return nil
}