package investment

import (
	"math"

	asset "github.com/andreposman/magic-number/internal/asset/model"
)

//CalculateGoals ...
func CalculateGoals(assetPrice float64, assetYieldAvarage24M float64, desiredMonthlyIncome float64) *asset.Goals {
	goals := new(asset.Goals)

	goals.DesiredMonthlyIncome = desiredMonthlyIncome
	goals.MagicNumber = math.Round(calculateMagicNumber(assetPrice, assetYieldAvarage24M))
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
