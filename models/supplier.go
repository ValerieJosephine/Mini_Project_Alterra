package models

//>tipe data supplier
type Supplier struct {
	IdSupplier   int    `json:"id supplier"`
	SupplierName string `json:"supplier name"`
	Contact      string `json:"email"`
	Address      string `json:"address"`
	Kurir        int    `json:"kurir"`
	UpdatedAt    int    `json:"updated"`
}
