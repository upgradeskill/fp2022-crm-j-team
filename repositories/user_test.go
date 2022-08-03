package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

var userModel = models.User{
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

	user, err := userRepository.Create(&userModel)
	newUser, _ := userRepository.Get(user.ID)
	userModel.ID = user.ID

	assert.Equal(t, userModel.Name, newUser.Name, "Name on retur Create user and get by id must match")
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

	userModel.Email = "wahyu@gmail.com"
	user, err := userRepository.Update(&userModel)

	assert.Equal(t, userModel.Email, user.Email, "Email must match")
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

	user, err := userRepository.Get(userModel.ID)

	assert.Equal(t, userModel.Email, user.Email, "Email must match")
	assert.Equal(t, "", err.Type, "Error Type must blank")
}

/**
* ==========================================
* Test Repository Login User
*===========================================
 */
func TestSuccessGetLogin(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	user, err := userRepository.Login(userModel.Email)

	assert.Equal(t, userModel.Email, user.Email, "Email must match")
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

	_, err := userRepository.GetAll(&schemas.User{})

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
	user, _ := userRepository.CheckEmailExistOnUpdate(newEmail, userModel.ID)

	assert.Equal(t, "", user.Email, "Email not registered yet")
}

func TestCheckEmailExistOnUpdate(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	newEmail := "wahyu@gmail.com"
	user, _ := userRepository.CheckEmailExistOnUpdate(newEmail, "123456")

	assert.Equal(t, newEmail, user.Email, "Email already registered")
}

/**
* ==========================================
* Test Repository Delete User
*===========================================
 */
func TestSuccessDelete(t *testing.T) {
	db := helpers.SetupDatabaseTesting()
	userRepository := NewRepositoryUser(db)

	_, errDelete := userRepository.Delete(&userModel)
	user, _ := userRepository.Get(userModel.ID)

	assert.Equal(t, "", errDelete.Type, "Error Type must empty")
	assert.Equal(t, "", user.ID, "User ID must empty")
}
