package service

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/asset/repository"
)

//GetAsset ...
func GetAsset(req *assetModel.Request) *assetModel.Asset {
	return buildAsset(req)
}

//buildAsset finds and calculates the asset data
func buildAsset(req *assetModel.Request) *assetModel.Asset {
	assetRaw := repository.FindAsset(req)
	calculationData := repository.AssetCalculation(assetRaw)
	assetCalculated := calculateInvestmentGoals(calculationData)
	asset := repository.GetAsset(assetRaw, assetCalculated)

	return asset
}

//calculateInvestmentGoals ...
func calculateInvestmentGoals(calculationData *assetModel.AssetDataFloat) *assetModel.AssetDataFloat {

	calculationData.MagicNumber = calculateMagicNumber(calculationData.Price, calculationData.YieldAvarage24M)
	calculationData.CapitalSnowBallEffect = calculateCapitalInvestedSnowball(calculationData.Price, calculationData.MagicNumber)
	calculationData.CapitalDesiredMonthlyIncome = calculateDesiredMonthlyIncome(calculationData.Price, calculationData.YieldAvarage24M, calculationData.DesiredMonthlyIncome)
	calculationData.AssetQuantityDesiredIncome = calculateQuantityDesiredIncome(calculationData.Price, calculationData.CapitalDesiredMonthlyIncome)
	return calculationData
}

//CalculateMagicNumber  calculates the amount of assets need to buy the asset itself with the dividends
// MagicNumber: assetPrice / assetYield = assetQuantity
func calculateMagicNumber(assetPrice float64, assetYieldAvarage24M float64) float64 {
	if assetPrice <= 0 || assetYieldAvarage24M <= 0 {
		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Price/Yield must be greater than zero")
		os.Exit(-1)
	}

	return float64(int64(math.Round(assetPrice / assetYieldAvarage24M)))
}

//CalculateCapitalInvestedSnowball calculates amount of money needed for the snowball effect
func calculateCapitalInvestedSnowball(assetPrice float64, magicNumber float64) float64 {
	if assetPrice <= 0 || magicNumber <= 0 {
		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Price/Magic Number must be greater than zero")
		os.Exit(-1)
	}
	return assetPrice * magicNumber
}

//calculateQuantityDesiredIncome
func calculateQuantityDesiredIncome(assetPrice float64, capitalDesiredMonthlyIncome float64) float64 {
	if assetPrice <= 0 || capitalDesiredMonthlyIncome <= 0 {
		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Price/Capital for Desired Monthly Income must be greater than zero")
		os.Exit(-1)
	}

	return (capitalDesiredMonthlyIncome / assetPrice)
}

//CalculateDesiredMonthlyIncome calculates the amount of capital needed to reach the desired monthly income from dividends
// ( (desired income/yield) * price)
func calculateDesiredMonthlyIncome(assetPrice float64, assetYieldAvarage24M float64, desiredMonthlyIncome float64) float64 {
	if assetPrice <= 0 || assetYieldAvarage24M <= 0 || desiredMonthlyIncome <= 0 {
		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Price/Yield/Desired Monthly Income must be greater than zero")
		os.Exit(-1)
	}
	return (desiredMonthlyIncome / assetYieldAvarage24M) * assetPrice
}

//BuildJSON ...
func BuildJSON(asset *assetModel.Asset) []byte {
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(assetJSON))

	return assetJSON
}
