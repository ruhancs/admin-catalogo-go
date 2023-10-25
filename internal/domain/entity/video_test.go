package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var categoriesIDs []string

func TestNewVideo(t *testing.T) {
	video,err := NewVideo("video1","",2004,1.30,categoriesIDs)

	assert.Nil(t,err)
	assert.NotNil(t,video)
	assert.NotNil(t,video.ID)
}

func TestChangeTitle(t *testing.T) {
	video,err := NewVideo("video1","",2004,1.30,categoriesIDs)
	assert.Nil(t,err)
	
	err = video.ChangeTitle("title updated")
	assert.Nil(t,err)
	assert.Equal(t,video.Title,"title updated")
}

func TestChangeInvalidTitle(t *testing.T) {
	video,err := NewVideo("video1","",2004,1.30,categoriesIDs)
	assert.Nil(t,err)
	
	err = video.ChangeTitle("")
	assert.NotNil(t,err)
	assert.Equal(t,video.Title,"video1")
}

func TestChangeDescription(t *testing.T) {
	video,err := NewVideo("video1","",2004,1.30,categoriesIDs)
	assert.Nil(t,err)

	video.ChangeDescription("desc1")
	assert.Equal(t,video.Description,"desc1")
}

func TestMarkPublished(t *testing.T) {
	video,err := NewVideo("video1","",2004,1.30,categoriesIDs)
	assert.Nil(t,err)

	video.MarkPublished()
	assert.Equal(t,video.IsPublished,true)
}