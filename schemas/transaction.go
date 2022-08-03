package schemas

type Transaction struct {
	OutletProductId string `json:"outlet_product_id" form:"outlet_product_id" validate:"required,uuid"`
	StaffId         string `json:"staff_id" form:"staff_id" validate:"required,uuid"`
	Qty             uint64 `json:"qty" form:"qty" validate:"required,numeric"`
}

type TransactionPage struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
