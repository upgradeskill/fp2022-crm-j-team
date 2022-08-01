package services

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

var createUserSchema = schemas.User{
	Name:     "Wahyu",
	Email:    "wahyuagung26@gmail.com",
	Password: "bismillah",
	Role:     "owner",
	OutletId: "1",
}

var createUserSecondSchema = schemas.User{
	Name:     "CRM",
	Email:    "crm@gmail.com",
	Password: "bismillah",
	Role:     "owner",
	OutletId: "1",
}

/**
* ==========================================
* Test Service Create User
*===========================================
 */
func TestSuccessCreate(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	user, err := service.Create(&createUserSchema)
	createUserSchema.ID = user.ID

	assert.Equal(t, user.Name, createUserSchema.Name, "Name must match")
	assert.Equal(t, "", err.Type, "Error Type must blank")
}

func TestFailedCreateEmailAlreadyExists(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	_, err := service.Create(&createUserSchema)

	assert.Equal(t, http.StatusUnprocessableEntity, err.Code, "Error Code must match")
}

/**
* ==========================================
* Test Service Create User
*===========================================
 */
func TestSuccessUpdate(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	createUserSchema.Email = "wahyu@gmail.com"
	user, err := service.Update(&createUserSchema)

	assert.Equal(t, user.Email, createUserSchema.Email, "Email must change and match")
	assert.Equal(t, "", err.Type, "Error Type must blank")
}

func TestFailedUpdateUserNotFound(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	createUserSchema.ID = "12345"
	_, err := service.Update(&createUserSchema)

	assert.Equal(t, http.StatusNotFound, err.Code, "Error Code must match")
}

func TestFailedUpdateEmailIsExist(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	user, err := service.Create(&createUserSecondSchema)
	createUserSecondSchema.ID = user.ID
	createUserSecondSchema.Email = createUserSchema.Email

	updatedUser, err := service.Update(&createUserSecondSchema)

	assert.Equal(t, updatedUser.Email, createUserSchema.Email, "Updated email must match")
	assert.Equal(t, http.StatusUnprocessableEntity, err.Code, "Error Code must match")
}

/**
* ==========================================
* Test Service Delete User
*===========================================
 */
func TestFailedDeleteUserNotFound(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	createUserSchema.ID = "09876543321"
	_, err := service.Delete(&createUserSchema)

	assert.Equal(t, http.StatusNotFound, err.Code, "User not found")
}

func TestSuccessDelete(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	newUser := schemas.User{
		Name:     "CRM",
		Email:    "crm@gmail.com",
		Password: "bismillah",
		Role:     "owner",
		OutletId: "1",
	}
	user, _ := service.Create(&newUser)

	newUser.ID = user.ID
	_, errDelete := service.Delete(&newUser)

	getUser, _ := service.Get(&newUser)

	assert.Equal(t, "", errDelete.Type, "Error Type for delete must empty")
	assert.Equal(t, "", getUser.ID, "ID User must empty")
}

/**
* ==========================================
* Test Service Get All User
*===========================================
 */
func TestSuccessGetAllUsers(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	repo := repositories.NewRepositoryUser(db)
	service := NewServiceUser(repo)

	_, err := service.GetAll()

	assert.Equal(t, "", err.Type, "Error Type must empty")
}
