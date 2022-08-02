package helpers

type UserLogin struct {
	ID       string
	Email    string
	Role     string
	OutletId string
}

func SessionUser() *UserLogin {
	user := UserLogin{
		ID:    "b21e233d-f73e-4c87-b994-d4bc74038dec",
		Email: "wahyu@majoo.id",
		Role:  "owner",
		OutletId: "fa326580-6bdf-44c7-b2fb-8492d5682aea	",
	}

	return &user
}
