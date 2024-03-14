package model

type Car struct {
	ID           string `json:"uuid"`
	Year         string `json:"year"`
	Make         string `json:"make"`
	Model        string `json:"model"`
	Trim         string `json:"trim"`
	Body         string `json:"body"`
	Transmission string `json:"trasmission"`
	State        string `json:"state"`
	Condition    string `json:"condition"`
	Odometer     string `json:"odometer"`
	Color        string `json:"color"`
	Interior     string `json:"interior"`
	Seller       string `json:"seller"`
	Mmr          string `json:"mmr"`
	SellingPrice string `json:"selling_price"`
	SaleDate     string `json:"sale_date"`
}
