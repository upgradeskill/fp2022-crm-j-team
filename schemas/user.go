package schemas

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"passowrd"`
	Role     string `json:"role"`
	OutletId string `json:"outlet_id"`
}
