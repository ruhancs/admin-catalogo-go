package usecase

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/entity"
	"admin-catalogo-go/pkg/events"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Insert(ctx context.Context, category *entity.Category) error {
	//args := r.Called(category)
	return nil
}

func (r *repositoryMock) Update(ctx context.Context, category *entity.Category) error {
	args := r.Called(category)
	return args.Error(0)
}

func (r *repositoryMock) ListCategory(ctx context.Context, params dto.ListCategoryInputDto) ([]entity.Category,error) {
	//args := r.Called()
	return nil, nil
}

func (r *repositoryMock) FindByID(ctx context.Context, id string) (entity.Category, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return entity.Category{}, args.Error(1)
	}
	return args.Get(0).(entity.Category), args.Error(1)
}

func (r *repositoryMock) Delete(ctx context.Context, id string) error {
	args := r.Called(id)
	return args.Error(0)
}

var repository = new(repositoryMock)

type categoryCreatedEventMock struct {
	mock.Mock
}

func (c *categoryCreatedEventMock) GetName() string {
	return "CategoryCreated"
}
func (c *categoryCreatedEventMock) GetDateTime() time.Time {
	return time.Now()
}
func (c *categoryCreatedEventMock) GetPayload() interface{} {
	return dto.CreateCategoryOutputDto{
		ID: "ashgey32r87gasdef",
		Name: "category1",
		Description: "description1",
		IsActive: true,
		CreatedAt: time.Now(),
	}
}
func (c *categoryCreatedEventMock) SetPayload(payload interface{}) {}
var categoryCreatedEvent = new(categoryCreatedEventMock)

type eventDispatcherMock struct {
	mock.Mock
}
func(e *eventDispatcherMock) Register(eventName string, handler events.EventHandlerInterface) error{
	return nil
}
func(e *eventDispatcherMock) Dispatch(event events.EventInterface) error{
	return nil
}
func(e *eventDispatcherMock) Remove(eventName string, handler events.EventHandlerInterface) error{
	return nil
}
func(e *eventDispatcherMock) Has(eventName string, handler events.EventHandlerInterface) bool{
	return true
}
func(e *eventDispatcherMock) Clear() error {
	return nil
}
var dispatcher = new(eventDispatcherMock)


func TestCreateCategoryUseCase(t *testing.T) {
	usecase := NewCreateCategoryUseCase(repository,categoryCreatedEvent,dispatcher)
	inputDto := dto.CreateCategoryInputDto{
		Name: "test",
		Description: "testing",
	}

	outputDto,err := usecase.Execute(context.Background(),inputDto)
	
	assert.NotNil(t,usecase)
	assert.Nil(t,err)
	assert.Equal(t,outputDto.Name,"test")
	assert.Equal(t,outputDto.Description,"testing")
	assert.Equal(t,outputDto.IsActive,true)
}