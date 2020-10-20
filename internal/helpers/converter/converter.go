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
	a.Asset.YieldAvarage24M = strconv.FormatFloat(assetNumber.Asset.YieldAvarage24M, 'f', 5, 64)
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
	yieldAvarage24M, err := strconv.ParseFloat(assetString.YieldAvarage24M, 5)
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
	// magicNumber, err := strconv.ParseFloat(assetString.Goals.MagicNumber, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// capitalSnowBallEffect, err := strconv.ParseFloat(assetString.Goals.CapitalSnowBallEffect, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// capitalDesiredMonthlyIncome, err := strconv.ParseFloat(assetString.Goals.CapitalDesiredMonthlyIncome, 2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

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
	// a.InvestmentGoals.MagicNumber = magicNumber
	// a.InvestmentGoals.CapitalSnowBallEffect = capitalSnowBallEffect
	// a.InvestmentGoals.CapitalDesiredMonthlyIncome = capitalDesiredMonthlyIncome

	return a
}
