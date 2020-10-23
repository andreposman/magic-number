package converter

import (
	"fmt"
	"strconv"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
)

//ToString converts type Asset and Goal to strings
func ToString(assetNumber *assetModel.ToNumberConverted) *assetModel.ToStringConverted {
	a := new(assetModel.ToStringConverted)

	a.Asset.Symbol = assetNumber.Asset.Symbol
	a.Asset.Name = assetNumber.Asset.Name
	a.Asset.Price = strconv.FormatFloat(assetNumber.Asset.Price, 'f', 2, 64)
	a.Asset.YieldAverage24M = strconv.FormatFloat(assetNumber.Asset.YieldAvarage24M, 'f', -1, 64)
	a.Asset.DividendYield = strconv.FormatFloat(assetNumber.Asset.DividendYield, 'f', 2, 64)
	a.Asset.MinPrice52Week = strconv.FormatFloat(assetNumber.Asset.MinPrice52Week, 'f', 2, 64)
	a.Asset.MaxPrice52Week = strconv.FormatFloat(assetNumber.Asset.MaxPrice52Week, 'f', 2, 64)
	a.Asset.PerformanceLast12M = assetNumber.Asset.PerformanceLast12M
	a.Asset.PerformanceThisMonth = assetNumber.Asset.PerformanceThisMonth
	a.Investment.DesiredMonthlyIncome = strconv.FormatFloat(assetNumber.InvestmentGoals.DesiredMonthlyIncome, 'f', 2, 64)
	a.Investment.MagicNumber = strconv.FormatFloat(assetNumber.InvestmentGoals.MagicNumber, 'f', 0, 64)
	a.Investment.CapitalSnowBallEffect = strconv.FormatFloat(assetNumber.InvestmentGoals.CapitalSnowBallEffect, 'f', 2, 64)
	a.Investment.CapitalDesiredMonthlyIncome = strconv.FormatFloat(assetNumber.InvestmentGoals.CapitalDesiredMonthlyIncome, 'f', 2, 64)

	return a
}

//ToFloat converts type Asset and Goal to strings
func ToFloat(assetString *assetModel.Model) *assetModel.ToNumberConverted {
	a := new(assetModel.ToNumberConverted)

	price, err := strconv.ParseFloat(assetString.Price, 2)
	if err != nil {
		fmt.Println(err)
	}
	yieldAvarage24M, err := strconv.ParseFloat(assetString.YieldAverage24M, 5)
	if err != nil {
		fmt.Println(err)
	}
	dividendYield, err := strconv.ParseFloat(assetString.DividendYield, 2)
	if err != nil {
		fmt.Println(err)
	}
	minPrice52Week, err := strconv.ParseFloat(assetString.MinPrice52Week, 2)
	if err != nil {
		fmt.Println(err)
	}
	maxPrice52Week, err := strconv.ParseFloat(assetString.MaxPrice52Week, 2)
	if err != nil {
		fmt.Println(err)
	}
	desiredMonthlyIncome, err := strconv.ParseFloat(assetString.Goals.DesiredMonthlyIncome, 2)
	if err != nil {
		fmt.Println(err)
	}

	a.Asset.Symbol = assetString.Symbol
	a.Asset.Name = assetString.Name
	a.Asset.Price = price
	a.Asset.YieldAvarage24M = yieldAvarage24M
	a.Asset.DividendYield = dividendYield
	a.Asset.MinPrice52Week = minPrice52Week
	a.Asset.MaxPrice52Week = maxPrice52Week
	a.Asset.PerformanceLast12M = assetString.PerformanceLast12M
	a.Asset.PerformanceThisMonth = assetString.PerformanceThisMonth
	a.InvestmentGoals.DesiredMonthlyIncome = desiredMonthlyIncome

	return a
}

// ToExpectedString ...
func ToExpectedString(value float64) string {
	valueString := strconv.FormatFloat(value, 'f', 2, 64)

	return valueString
}

// ToYieldAvarageString ...
func ToYieldAvarageString(value float64) string {
	valueString := strconv.FormatFloat(value, 'f', 5, 64)

	return valueString
}

// ToMagicNumberString ...
func ToMagicNumberString(value float64) string {
	valueString := strconv.FormatFloat(value, 'f', 0, 64)

	return valueString
}

// ToExpectedFloat ...
func ToExpectedFloat(text string) float64 {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Print("\nToExpectedFloat\n")
		fmt.Println(err)
	}

	return value
}

// ToYieldAverageFloat ...
func ToYieldAverageFloat(text string) float64 {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Print("\nToYieldAverageFloat\n")
		fmt.Println(err)
	}

	return value
}
