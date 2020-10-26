package asset

//Request is the data inputted by the user
type Request struct {
	AssetSymbol          string
	DesiredMonthlyIncome string
}

//Elements is the model of the html elements scrapped by the crawler
type Elements struct {
	Symbol               string
	Name                 string
	Price                string
	YieldAverage24M      string
	DividendYield        string
	MinPrice52Week       string
	MaxPrice52Week       string
	PerformanceLast12M   string
	PerformanceThisMonth string
}

//Asset is the model that contains the value of the asset
type Asset struct {
	Symbol               string `json:"symbol"`
	Name                 string `json:"name"`
	Price                string `json:"price"`
	YieldAverage24M      string `json:"yieldAverage24m"`
	DividendYield        string `json:"dividendYield"`
	MinPrice52Week       string `json:"minPrice52wk"`
	MaxPrice52Week       string `json:"maxPrice52wk"`
	PerformanceLast12M   string `json:"performanceLast12m"`
	PerformanceThisMonth string `json:"performanceThisMonth"`
	Goals                Goals  `json:"goals"`
}

//Goals has the information for the investor about capital needed to achieve some goals
type Goals struct {
	MagicNumber                 string `json:"magicNumber"`
	CapitalSnowBallEffect       string `json:"capitalSnowBallEffect"`
	DesiredMonthlyIncome        string `json:"desiredMonthlyIncome"`
	AssetQuantityDesiredIncome  string `json:"assetQuantityDesiredIncome"`
	CapitalDesiredMonthlyIncome string `json:"capitalDesiredMonthlyIncome"`
}

//AssetDataFloat is the strongly typed model
type AssetDataFloat struct {
	Price                       float64
	YieldAvarage24M             float64
	MagicNumber                 float64
	CapitalSnowBallEffect       float64
	DesiredMonthlyIncome        float64
	AssetQuantityDesiredIncome  float64
	CapitalDesiredMonthlyIncome float64
}
