package schemas

type Category struct {
	ID          string `json:"id" validate:"uuid"`
	Name        string `json:"name" validate:"required,lowercase"`
	Description string `json:"description"`
	Page        int    `json:"page"`
}
