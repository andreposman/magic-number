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

//AssetNumber is the strongly typed model
type AssetNumber struct {
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

//Goals has the information for the investor about capital needed to achieve some goals
type Goals struct {
	MagicNumber                 float64
	CapitalSnowBallEffect       float64
	DesiredMonthlyIncome        float64
	CapitalDesiredMonthlyIncome float64
}

//GoalsString has the information for the investor about capital needed to achieve some goals
type GoalsString struct {
	MagicNumber                 string
	CapitalSnowBallEffect       string
	DesiredMonthlyIncome        string
	CapitalDesiredMonthlyIncome string
}

type RequestAsset struct {
	AssetSymbol          string
	DesiredMonthlyIncome string
}

//ToStringConverted is the result of float64 to string
type ToStringConverted struct {
	AssetString           AssetString
	InvestmentGoalsString GoalsString
}

//ToNumberConverted is the result of string to float64
type ToNumberConverted struct {
	Asset           AssetNumber
	InvestmentGoals Goals
}
