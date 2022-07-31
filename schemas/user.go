package schemas

type User struct {
	ID       string `json:"id" validate:"uuid"`
	Name     string `json:"name" validate:"required,lowercase"`
	Email    string `json:"email" validate:"required,lowercase"`
	Password string `json:"passowrd"`
	Role     string `json:"role"`
	OutletId string `json:"outlet_id"`
}
