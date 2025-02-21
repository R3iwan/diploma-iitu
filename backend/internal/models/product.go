package models

type Product struct {
	CompanyID   int    `json:"company_id"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	ProductType string `json:"product_type"`
	ProductDesc string `json:"product_desc"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
