package service

import (
	"encoding/json"
	"fmt"
	"math"

	asset "github.com/andreposman/magic-number/internal/asset/model"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/crawler/service"
	"github.com/andreposman/magic-number/internal/helpers/converter"
)

//GetAsset ...
func GetAsset(req *asset.Request) *asset.ToStringConverted {
	a := service.FindAsset(req)
	asset := CalculateGoals(a)

	return asset
}

//CalculateGoals ...
func CalculateGoals(asset *asset.Model) *assetModel.ToStringConverted {
	a := converter.ToFloat(asset)

	a.InvestmentGoals.MagicNumber = calculateMagicNumber(a.Asset.Price, a.Asset.YieldAvarage24M)
	a.InvestmentGoals.CapitalSnowBallEffect = calculateCapitalInvestedSnowball(a.Asset.Price, a.InvestmentGoals.MagicNumber)
	a.InvestmentGoals.CapitalDesiredMonthlyIncome = calculateDesiredMonthlyIncome(a.Asset.Price, a.Asset.YieldAvarage24M, a.InvestmentGoals.DesiredMonthlyIncome)

	aString := converter.ToString(a)

	return aString
}

/*CalculateMagicNumber  calculates the amount of assets need to buy the asset itself with the dividends
MagicNumber: assetPrice / assetYield = assetQuantity */
func calculateMagicNumber(assetPrice float64, assetYieldAvarage24M float64) float64 {
	return math.Round(assetPrice / assetYieldAvarage24M)
}

//CalculateCapitalInvestedSnowball calculates amount of money needed for the snowball effect
func calculateCapitalInvestedSnowball(assetPrice float64, magicNumber float64) float64 {
	return assetPrice * magicNumber
}

//CalculateDesiredMonthlyIncome calculates the amount of capital needed to reach the desired monthly income from dividends ( (desired income/yield) * price)
func calculateDesiredMonthlyIncome(assetPrice float64, assetYieldAvarage24M float64, desiredMonthlyIncome float64) float64 {
	return (desiredMonthlyIncome / assetYieldAvarage24M) * assetPrice
}

//BuildJSON ...
func BuildJSON(req *asset.ToStringConverted) []byte {

	a := &assetModel.ToStringConverted{
		Asset:      req.Asset,
		Investment: req.Investment,
	}

	assetJSON, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(assetJSON))

	return assetJSON
}
