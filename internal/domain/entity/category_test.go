package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	category,err := NewCategory("C123","D1")

	assert.Nil(t,err)
	assert.NotNil(t,category)
	assert.NotNil(t,category.ID)
	assert.Equal(t,category.Name,"C123")
	assert.Equal(t,category.Description,"D1")
	assert.Equal(t,category.IsActive,true)
}

func TestCategoryDeActive(t *testing.T) {
	category,err := NewCategory("C123","D1")

	category.DeActive()

	assert.Nil(t,err)
	assert.NotNil(t,category)
	assert.Equal(t,category.Name,"C123")
	assert.Equal(t,category.Description,"D1")
	assert.Equal(t,category.IsActive,false)
}

func TestCategoryUpdate(t *testing.T) {
	category,err := NewCategory("C123","D1")

	category.Update("updated", "updated")

	assert.Nil(t,err)
	assert.NotNil(t,category)
	assert.Equal(t,category.Name,"updated")
	assert.Equal(t,category.Description,"updated")
	assert.Equal(t,category.IsActive,true)
}

func TestCategoryNameRequired(t *testing.T) {
	category,err := NewCategory("","D1")

	assert.Nil(t,category)
	assert.NotNil(t,err)
	assert.Error(t,err, "Expected value not to be nil")
}

func TestCategoryDescriptionRequired(t *testing.T) {
	category,err := NewCategory("name","")

	assert.Nil(t,category)
	assert.NotNil(t,err)
	assert.Error(t,err, "Expected value not to be nil")
}