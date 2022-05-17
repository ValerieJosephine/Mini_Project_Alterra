package models

//>tipe data logbook stock in
type StockIn struct {
	IdProduct     int    `json:"id product"`
	IdProductType int    `json:"id product type"`
	Description   string `json:"desc"`
	Quantity      string `json:"quantity"`
	StaffName     string `json:"staff name"`
	Price         int    `json:"price"`
	Reorderat     int    `json:"reorder"`
	IdSupplier    int    `json:"id supplier"`
	UpdatedAt     int    `json:"updated"`
}
