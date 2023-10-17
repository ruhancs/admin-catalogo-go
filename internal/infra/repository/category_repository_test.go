package repository

import (
	"admin-catalogo-go/internal/application/dto"
	"admin-catalogo-go/internal/domain/entity"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func createTestCategory() *entity.Category{
	category,_ := entity.NewCategory("test", "testing")
	return category 
}

func initRepository() *CategoryRepository{
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Println("Error loading .env")
	}
	
	dbDriver := os.Getenv("DB_DRIVER_TEST")
	dbSource := os.Getenv("DB_SOURCE_TEST")
	
	testDB,err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println(err)
		log.Fatal("cannot connect to db")
	}

	repository := NewCategoryRepository(testDB)
	return repository
}


func TestCreateCategory(t *testing.T) {
	
	var repository = initRepository()
	
	category := createTestCategory()

	err := repository.Insert(context.Background(),category)

	assert.Nil(t,err)
}

func TestFindByIDCategory(t *testing.T) {
	var repository = initRepository()
	
	createdCategory := createTestCategory()

	err := repository.Insert(context.Background(),createdCategory)
	assert.Nil(t,err)

	categoryFounded,err := repository.FindByID(context.Background(),createdCategory.ID)

	assert.Nil(t,err)
	assert.Equal(t,categoryFounded.Name,createdCategory.Name)
	assert.Equal(t,categoryFounded.Description,createdCategory.Description)
	assert.Equal(t,categoryFounded.IsActive,createdCategory.IsActive)

	categoryFounded,err = repository.FindByID(context.Background(),"123456789123456")
	
	assert.NotNil(t,err)
}

func TestListCategory(t *testing.T) {
	var repository = initRepository()
	
	//createdCategory := createTestCategory()
	//createdCategory2 := createTestCategory()
	//createdCategory3 := createTestCategory()

	//err := repository.Insert(context.Background(),createdCategory)
	//err = repository.Insert(context.Background(),createdCategory2)
	//err = repository.Insert(context.Background(),createdCategory3)
	//assert.Nil(t,err)

	listCategoryDto := dto.ListCategoryInputDto{
		PerPage: 10,
	}
	categories,err := repository.ListCategory(context.Background(),listCategoryDto)

	assert.Nil(t,err)
	assert.Len(t,categories,3)
}

func TestDeleteCategory(t *testing.T) {
	var repository = initRepository()
	
	createdCategory := createTestCategory()

	err := repository.Insert(context.Background(),createdCategory)
	assert.Nil(t,err)

	err = repository.Delete(context.Background(),createdCategory.ID)

	assert.Nil(t,err)
}
