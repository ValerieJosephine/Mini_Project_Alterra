package models

//>tipe data product_type
type ProductType struct {
	IdProductType int    `json:"id"`
	RollcaseType  string `json:"rollcase"`
	BrushcaseType string `json:"brushcase"`
	MinibagType   string `json:"minibag"`
	UpdatedAt     int    `json:"updated"`
}
