package schemas

type Product struct {
	ID         string `json:"id" validate:"uuid"`
	Name       string `json:"name" validate:"required,lowercase"`
	SKU        uint64 `json:"sku" validate:"required,numeric"`
	CategoryId uint64 `json:"category_id" validate:"required,uuid"`
}
