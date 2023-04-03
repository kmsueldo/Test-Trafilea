package models

type Cart struct {
	CartID   string    `json:"CartID"`
	UserID   string    `json:"UserID"`
	Products []Product `json:"Products"`
}

type Product struct {
	ProductID string  `json:"ProductID"`
	Name      string  `json:"Name"`
	Category  string  `json:"Category"`
	Price     float32 `json:"Price"`
	Quantity  int     `json:"Quantity"`
}

type Order struct {
	CartID string `json:"CartID"`
	Totals Totals `json:"Totals"`
}

type Totals struct {
	Products        []Product `json:"Products"`
	AmountShipping  int       `json:"Amount_shipping"`
	PercentDiscount int       `json:"Percent_discount"`
	AmountOrder     float32   `json:"Amount_total_order"`
}
