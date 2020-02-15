package model

import "time"

type Supplier struct {
	ID int `form:"id"`
	SupplierName string `form:"name"`
	Tel string `form:"tel"`
	Address string `form:"address"`
	CreatedAt time.Time
}
