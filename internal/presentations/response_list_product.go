package presentations

type ResponseListingProduct struct {
	Code     int            `json:"code"`
	Message  string         `json:"message"`
	Data     []DataProducts `json:"data,omitempty"`
	MetaData MetaData       `json:"meta_data,omitempty"`
}

type DataProducts struct {
	Id                int               `json:"id"`
	SellerInformation SellerInformation `json:"seller_information"`
	ProductName       string            `json:"product_name"`
	Price             float64           `json:"price"`
	Description       string            `json:"description"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
}

type SellerInformation struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}
