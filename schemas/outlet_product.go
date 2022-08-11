package schemas

type OutletProduct struct {
	ID        string  `json:"id" form:"id"`
	OutletId  string  `json:"outlet_id" form:"outlet_id"`
	ProductId string  `json:"product_id" form:"product_id"`
	Price     float64 `json:"price" form:"price"`
	Stock     int     `json:"stock" form:"stock"`
	Page      int     `json:"page"`
	PageSize  int     `json:"page_size"`
}
