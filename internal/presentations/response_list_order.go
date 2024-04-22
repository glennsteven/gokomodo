package presentations

type ResponseListingOrderBuyer struct {
	Code     int              `json:"code"`
	Message  string           `json:"message"`
	Data     []DataOrderBuyer `json:"data,omitempty"`
	MetaData MetaData         `json:"meta_data,omitempty"`
}

type DataOrderBuyer struct {
	Id               int              `json:"id"`
	BuyerInformation BuyerInformation `json:"buyer_information"`
	DataProductBuyer DataProductBuyer `json:"data_product_buyer"`
	Status           int              `json:"status"`
	CreatedAt        string           `json:"created_at"`
	UpdatedAt        string           `json:"updated_at,omitempty"`
}

type DataProductBuyer struct {
	InformationSeller InformationSeller `json:"information_seller"`
	ProductName       string            `json:"product_name"`
	TotalPrice        float64           `json:"total_price"`
	Description       string            `json:"description"`
	Quantity          int               `json:"quantity"`
}

type InformationSeller struct {
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}
