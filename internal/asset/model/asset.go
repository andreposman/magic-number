package asset

//Model is the model of the asset elements scrapped by the crawler
type Elements struct {
	Symbol               string `json:"symbol"`
	Name                 string `json:"name"`
	Price                string `json:"price"`
	YieldAverage24M      string `json:"yieldAverage24m"`
	DividendYield        string `json:"dividendYield"`
	MinPrice52Week       string `json:"minPrice52wk"`
	MaxPrice52Week       string `json:"maxPrice52wk"`
	PerformanceLast12M   string `json:"performanceLast12m"`
	PerformanceThisMonth string `json:"performanceThisMonth"`
}

//Model is the model of the asset elements scrapped by the crawler
type Model struct {
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

//ModelFloat is the strongly typed model
type ModelFloat struct {
	Symbol               string
	Name                 string
	Price                float64
	YieldAvarage24M      float64
	DividendYield        float64
	MinPrice52Week       float64
	MaxPrice52Week       float64
	PerformanceLast12M   string
	PerformanceThisMonth string
	Goals                GoalsNumber
}

//GoalsNumber has the information for the investor about capital needed to achieve some goals
type GoalsNumber struct {
	MagicNumber                 float64
	CapitalSnowBallEffect       float64
	DesiredMonthlyIncome        float64
	CapitalDesiredMonthlyIncome float64
}

//Goals has the information for the investor about capital needed to achieve some goals
type Goals struct {
	MagicNumber                 string `json:"magicNumber"`
	CapitalSnowBallEffect       string `json:"capitalSnowBallEffect"`
	DesiredMonthlyIncome        string `json:"desiredMonthlyIncome"`
	CapitalDesiredMonthlyIncome string `json:"capitalDesiredMonthlyIncome"`
}

type Request struct {
	AssetSymbol          string
	DesiredMonthlyIncome string
}

//ToStringConverted is the result of float64 to string
type ToStringConverted struct {
	Asset      Model `json:"asset"`
	Investment Goals `json:"investment"`
}

//ToNumberConverted is the result of string to float64
type ToNumberConverted struct {
	Asset           ModelFloat
	InvestmentGoals GoalsNumber
}

type DataTable struct {
	Symbol                      string
	Name                        string
	Price                       string
	YieldAvarage24M             string
	DividendYield               string
	MinPrice52Week              string
	MaxPrice52Week              string
	PerformanceLast12M          string
	PerformanceThisMonth        string
	DesiredMonthlyIncome        string
	MagicNumber                 string
	CapitalSnowBallEffect       string
	CapitalDesiredMonthlyIncome string
}
