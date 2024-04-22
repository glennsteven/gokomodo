package presentations

type ResponseListingOrderDetails struct {
	Code     int                `json:"code"`
	Message  string             `json:"message"`
	Data     []DataOrderDetails `json:"data,omitempty"`
	MetaData MetaData           `json:"meta_data,omitempty"`
}

type DataOrderDetails struct {
	Id                int               `json:"id"`
	SellerInformation SellerInformation `json:"seller_information"`
	BuyerInformation  BuyerInformation  `json:"buyer_information"`
	DataProductSeller DataProductSeller `json:"data_product_seller"`
	Status            int               `json:"status"`
	CreatedAt         string            `json:"created_at"`
	UpdatedAt         string            `json:"updated_at"`
}

type DataProductSeller struct {
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}

type BuyerInformation struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	Email    string `json:"email,omitempty"`
}
