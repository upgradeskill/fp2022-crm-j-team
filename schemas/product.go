package schemas

type Product struct {
	ID         string `json:"id" form:"id" validate:"uuid"`
	Name       string `json:"name" form:"name" validate:"required,lowercase"`
	SKU        uint64 `json:"sku" form:"sku" validate:"required,numeric"`
	CategoryId string `json:"category_id" form:"category_id" validate:"required,uuid"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}
