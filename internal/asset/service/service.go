package service

import (
	"math"
	"strconv"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/crawler"
)

//CalculateGoals ...
func CalculateGoals(assetSymbol string, desiredMonthlyIncome float64) *assetModel.Goals {
	a := crawler.GetAsset(assetSymbol)
	goals := new(assetModel.Goals)

	goals.DesiredMonthlyIncome = desiredMonthlyIncome
	goals.MagicNumber = math.Round(calculateMagicNumber(a.Price, assetYieldAvarage24M))
	goals.CapitalSnowBallEffect = math.Round(calculateCapitalInvestedSnowball(assetPrice, goals.MagicNumber))
	goals.CapitalDesiredMonthlyIncome = float64(calculateDesiredMonthlyIncome(assetPrice, assetYieldAvarage24M, desiredMonthlyIncome))

	return goals
}

/*CalculateMagicNumber  calculates the amount of assets need to buy the asset itself with the dividends
MagicNumber: assetPrice / assetYield = assetQuantity */
func calculateMagicNumber(assetPrice float64, assetYieldAvarage24M float64) float64 {
	return (assetPrice / assetYieldAvarage24M)
}

//CalculateCapitalInvestedSnowball calculates amount of money needed for the snowball effect
func calculateCapitalInvestedSnowball(assetPrice float64, magicNumber float64) float64 {
	return (assetPrice * magicNumber)
}

//CalculateDesiredMonthlyIncome calculates the amount of capital needed to reach the desired monthly income from dividends ( (desired income/yield) * price)
func calculateDesiredMonthlyIncome(assetPrice float64, assetYieldAvarage24M float64, desiredMonthlyIncome float64) float64 {
	return (desiredMonthlyIncome / assetYieldAvarage24M) * assetPrice
}

//ConvertAssetToString converts type Asset and Goal to strings
func ConvertAssetToString(asset *assetModel.AssetNumber, goals *assetModel.Goals, desiredMonthlyIncome float64) *assetModel.ToStringConverted {
	a := new(assetModel.ToStringConverted)

	a.AssetString.Symbol = asset.Symbol
	a.AssetString.Name = asset.Name
	a.AssetString.Price = strconv.FormatFloat(asset.Price, 'f', 2, 64)
	a.AssetString.YieldAvarage24M = strconv.FormatFloat(asset.YieldAvarage24M, 'f', 5, 64)
	a.AssetString.DividendYield = strconv.FormatFloat(asset.Price, 'f', 2, 64)
	a.AssetString.MinPrice52Week = strconv.FormatFloat(asset.MinPrice52Week, 'f', 2, 64)
	a.AssetString.MaxPrice52Week = strconv.FormatFloat(asset.MaxPrice52Week, 'f', 2, 64)
	a.AssetString.PerformanceLast12M = asset.PerformanceLast12M
	a.AssetString.PerformanceThisMonth = asset.PerformanceThisMonth
	a.InvestmentGoalsString.DesiredMonthlyIncome = strconv.FormatFloat(desiredMonthlyIncome, 'f', 3, 64)
	a.InvestmentGoalsString.MagicNumber = strconv.FormatFloat(goals.MagicNumber, 'f', 0, 64)
	a.InvestmentGoalsString.CapitalSnowBallEffect = strconv.FormatFloat(goals.CapitalSnowBallEffect, 'f', 2, 64)
	a.InvestmentGoalsString.CapitalDesiredMonthlyIncome = strconv.FormatFloat(goals.CapitalDesiredMonthlyIncome, 'f', 5, 64)

	return a
}

//ConvertAssetToString converts type Asset and Goal to strings
func ConvertAssetToNumber(asset *assetModel.ToStringConverted) *assetModel.ToNumberConverted {
	a := new(assetModel.ToNumberConverted)

	a.Asset.Symbol = asset.AssetString.Symbol
	a.Asset.Name = asset.AssetString.Name
	a.Asset.Price = strconv.FormatFloat(asset.AssetString.Price, 'f', 2, 64)
	a.Asset.YieldAvarage24M = strconv.FormatFloat(asset.AssetString.YieldAvarage24M, 'f', 5, 64)
	a.Asset.DividendYield = strconv.FormatFloat(asset.AssetString.DividendYield, 'f', 2, 64)
	a.Asset.MinPrice52Week = strconv.FormatFloat(asset.AssetString.MinPrice52Week, 'f', 2, 64)
	a.Asset.MaxPrice52Week = strconv.FormatFloat(asset.AssetString.MaxPrice52Week, 'f', 2, 64)
	a.Asset.PerformanceLast12M = asset.PerformanceLast12M
	a.Asset.PerformanceThisMonth = asset.PerformanceThisMonth
	a.InvestmentGoal.DesiredMonthlyIncome = strconv.FormatFloat(asset.InvestmentGoalsString.DesiredMonthlyIncome, 'f', 3, 64)
	a.InvestmentGoal.MagicNumber = strconv.FormatFloat(asset.InvestmentGoalsString.MagicNumber, 'f', 2, 64)
	a.InvestmentGoal.CapitalSnowBallEffect = strconv.FormatFloat(asset.InvestmentGoalsString.CapitalSnowBallEffect, 'f', 2, 64)
	a.InvestmentGoal.CapitalDesiredMonthlyIncome = strconv.FormatFloat(goals.CapitalDesiredMonthlyIncome, 'f', 5, 64)

	return a
}
