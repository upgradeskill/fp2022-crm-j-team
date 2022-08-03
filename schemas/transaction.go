package schemas

type Transaction struct {
	OutletProductId string `json:"outlet_product_id" validate:"required,uuid"`
	StaffId         string `json:"staff_id" validate:"required,uuid"`
	Qty             uint64 `json:"qty" validate:"required,numeric"`
}

type TransactionPage struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
