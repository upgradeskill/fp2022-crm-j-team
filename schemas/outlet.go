package schemas

type Outlet struct {
	ID       string `json:"id" validate:"uuid"`
	Name     string `json:"name" validate:"required,lowercase"`
	Phone    int    `json:"phone"`
	Address  string `json:"description"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}
