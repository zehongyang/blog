package model

import "time"

type Bill struct {
	ID int `form:"id"`
	ProductName string `form:"product_name"`
	ProductUnit string `form:"product_unit"`
	ProductNum int `form:"product_num"`
	TotalAccount float64 `form:"total_count"`
	SupplierId int `form:"supplier_id"`
	Supplier Supplier
	IsPay int `form:"is_pay"`
	CreatedAt time.Time
}
