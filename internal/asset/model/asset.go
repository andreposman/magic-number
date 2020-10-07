package asset

//AssetElements is the model of the asset elements scrapped by the crawler
type AssetElements struct {
	Symbol               string //`json:"symbol,omitempty"`
	Name                 string //`json:"name,omitempty"`
	Price                string //`json:"price,omitempty"`
	YieldAvarage24M      string //`json:"yieldAvarage24M,omitempty"`
	DividendYield        string //`json:"dividendYield,omitempty"`
	MinPrice52Week       string //`json:"MinPrice52Week,omitempty"`
	MaxPrice52Week       string //`json:"MaxPrice52Week,omitempty"`
	PerformanceLast12M   string //`json:"PerformanceLast12M,omitempty"`
	PerformanceThisMonth string //`json:"PerformanceThisMonth,omitempty"`
}

//Asset is the strongly typed model
type Asset struct {
	Symbol               string
	Name                 string
	Price                float64
	YieldAvarage24M      float64
	DividendYield        float64
	MinPrice52Week       float64
	MaxPrice52Week       float64
	PerformanceLast12M   string
	PerformanceThisMonth string
}
