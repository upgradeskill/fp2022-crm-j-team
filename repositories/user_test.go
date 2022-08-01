package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
)

var createUserModel = models.User{
	Name:     "Wahyu",
	Email:    "wahyuagung26@gmail.com",
	Password: "bismillah",
	Role:     "owner",
	OutletId: "1",
}

/**
* ==========================================
* Test Repository Create User
*===========================================
 */
func TestSuccessInsert(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	user, err := userRepository.Create(&createUserModel)
	newUser, _ := userRepository.Get(user.ID)
	createUserModel.ID = user.ID

	assert.Equal(t, createUserModel.Name, newUser.Name, "Name on retur Create user and get by id must match")
	assert.Equal(t, "", err.Type, "Error Type must blank")
}

/**
* ==========================================
* Test Repository Update User
*===========================================
 */
func TestSuccessUpdate(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	createUserModel.Email = "wahyu@gmail.com"
	user, err := userRepository.Update(&createUserModel)

	assert.Equal(t, createUserModel.Email, user.Email, "Email must match")
	assert.Equal(t, "", err.Type, "Error Type must blank")
}

/**
* ==========================================
* Test Repository Get User
*===========================================
 */
func TestSuccessGetId(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	user, err := userRepository.Get(createUserModel.ID)

	assert.Equal(t, createUserModel.Email, user.Email, "Email must match")
	assert.Equal(t, "", err.Type, "Error Type must blank")
}

/**
* ==========================================
* Test Repository Get All User
*===========================================
 */
func TestSuccessGetAllUser(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	_, err := userRepository.GetAll()

	assert.Equal(t, "", err.Type, "Error Type must blank")
}

/**
* ==========================================
* Test Repository Check email exist on create
*===========================================
 */
func TestCheckEmailExistOnCreate(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	newEmail := "wahyuagung@gmail.com"
	user, _ := userRepository.CheckEmailExistOnCreate(newEmail)

	assert.Equal(t, "", user.Email, "Email not registered yet")
}

/**
* ==========================================
* Test Repository Check email exist on Update
*===========================================
 */
func TestCheckEmailNotExistOnUpdate(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	newEmail := "wahyu@gmail.com"
	user, _ := userRepository.CheckEmailExistOnUpdate(newEmail, createUserModel.ID)

	assert.Equal(t, "", user.Email, "Email not registered yet")
}

func TestCheckEmailExistOnUpdate(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	newEmail := "wahyu@gmail.com"
	user, _ := userRepository.CheckEmailExistOnUpdate(newEmail, "123456")

	assert.Equal(t, newEmail, user.Email, "Email already registered")
}
