package service

import (
	"encoding/json"
	"fmt"
	"math"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/crawler/service"
	"github.com/andreposman/magic-number/internal/helpers/converter"
)

//GetAsset ...
func GetAsset(req *assetModel.Request) *assetModel.Model {
	a := service.FindAsset(req)
	asset := CalculateInvestmentGoals(a)

	return asset
}

//CalculateInvestmentGoals ...
func CalculateInvestmentGoals(asset *assetModel.Model) *assetModel.Model {
	fmt.Print("\nPrice: ", asset.Price)
	price := converter.ToExpectedFloat(asset.Price)

	fmt.Print("\nyieldAverage24M: ", asset.YieldAverage24M)
	yieldAverage24M := converter.ToYieldAverageFloat(asset.YieldAverage24M)

	fmt.Print("\ndesiredMonthlyIncome: ", asset.Goals.DesiredMonthlyIncome)
	desiredMonthlyIncome := converter.ToExpectedFloat(asset.Goals.DesiredMonthlyIncome)

	asset.Goals.MagicNumber = converter.ToExpectedString(calculateMagicNumber(price, yieldAverage24M))
	asset.Goals.CapitalSnowBallEffect = converter.ToExpectedString(calculateCapitalInvestedSnowball(price, converter.ToExpectedFloat(asset.Goals.MagicNumber)))
	asset.Goals.CapitalDesiredMonthlyIncome = converter.ToExpectedString(calculateDesiredMonthlyIncome(price, yieldAverage24M, desiredMonthlyIncome))

	return asset
}

//CalculateMagicNumber  calculates the amount of assets need to buy the asset itself with the dividends
// MagicNumber: assetPrice / assetYield = assetQuantity
func calculateMagicNumber(assetPrice float64, assetYieldAvarage24M float64) float64 {
	return math.Round(assetPrice / assetYieldAvarage24M)
}

//CalculateCapitalInvestedSnowball calculates amount of money needed for the snowball effect
func calculateCapitalInvestedSnowball(assetPrice float64, magicNumber float64) float64 {
	return assetPrice * magicNumber
}

//CalculateDesiredMonthlyIncome calculates the amount of capital needed to reach the desired monthly income from dividends
// ( (desired income/yield) * price)
func calculateDesiredMonthlyIncome(assetPrice float64, assetYieldAvarage24M float64, desiredMonthlyIncome float64) float64 {
	return (desiredMonthlyIncome / assetYieldAvarage24M) * assetPrice
}

//BuildJSON ...
func BuildJSON(req *assetModel.Model) []byte {
	assetJSON, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(assetJSON))

	return assetJSON
}
