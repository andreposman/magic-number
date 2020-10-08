package asset

//AssetString is the model of the asset elements scrapped by the crawler
type AssetString struct {
	Symbol               string
	Name                 string
	Price                string
	YieldAvarage24M      string
	DividendYield        string
	MinPrice52Week       string
	MaxPrice52Week       string
	PerformanceLast12M   string
	PerformanceThisMonth string
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
