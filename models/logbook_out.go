package models

//>tipe data logbook stock out
type StockOut struct {
	IdProduct     int    `json:"id product"`
	IdProductType int    `json:"id product type"`
	Description   string `json:"desc"`
	Quantity      string `json:"quantity"`
	Price         int    `json:"price"`
	Reorderat     int    `json:"reorder"`
	IdSupplier    int    `json:"id supplier"`
	UpdatedAt     int    `json:"updated"`
}
