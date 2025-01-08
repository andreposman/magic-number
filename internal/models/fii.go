package models

type FII struct {
	ID            int     `json:"id"`
	Ticker        string  `json:"ticker"`
	FullName      string  `json:"company_name"`
	Price         float32 `json:"price"`
	DividendYield float32 `json:"dividend_yield"`
	Segment       string  `json:"segment"`
	Type          string  `json:"type"`
}

type FIIDetails struct {
	Title         string  `json:"title"`
	DividendYield float32 `json:"dividend_yield"`
	Segment       string  `json:"segment"`
	Type          string  `json:"type"`
}

type FIIDetailsResponse struct {
	Data []FIIDetails `json:"data"`
}
