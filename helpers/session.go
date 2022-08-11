package helpers

type UserLogin struct {
	ID       string
	Email    string
	Role     string
	OutletId string
}

var userLogin = UserLogin{}

func SessionUser() *UserLogin {
	user := UserLogin{
		ID:       userLogin.ID,
		Email:    userLogin.Email,
		Role:     userLogin.Role,
		OutletId: userLogin.OutletId,
	}

	return &user
}

func SetSession(id string, email string, role string, outlet_id string) *UserLogin {
	userLogin.ID = id
	userLogin.Email = email
	userLogin.Role = role
	userLogin.OutletId = outlet_id

	return &userLogin
}
