package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/upgradeskill/fp2022-crm-j-team/helper"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

/**
* ==========================================
* Repository Create User
*===========================================
 */
func TestSuccessInsert(t *testing.T) {
	db := helper.SetupDatabaseTesting()

	userSchema := schemas.User{
		Name:     "Wahyu",
		Email:    "wahyuagung26@gmail.com",
		Password: "bismillah",
		Role:     "owner",
		OutletId: "1",
	}

	userRepository := NewRepositoryUser(db)
	user, err := userRepository.Create(&userSchema)

	newUser, _ := userRepository.Get(user.ID)

	assert.Equal(t, userSchema.Name, newUser.Name, "Error code must blank")
	assert.Equal(t, "", err.Type, "Error code must blank")
}

func TestFailedInsertUser(t *testing.T) {
	db := helper.SetupDatabaseTesting()

	userSchema := schemas.User{
		Name:     "Wahyu",
		Email:    "wahyuagung26@gmail.com",
		Password: "bismillah",
		Role:     "owner",
		OutletId: "1",
	}

	userRepository := NewRepositoryUser(db)
	_, err := userRepository.Create(&userSchema)

	assert.Equal(t, "user already exists", err.Type, "User already exists")
}
