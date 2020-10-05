package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/andreposman/magic-number/internal/crawler"
)

func main() {

	assetSymbol := readAssetSymbolFromTerminal()
	var desiredMonthlyIncome float64 = 1000
	asset := crawler.BuildAsset(assetSymbol)

	assetQuantity := math.Round(calculateMagicNumber(asset.Price, asset.Yield))
	capitalInvested := math.Round(calculateCapitalInvestedSnowball(asset.Price, assetQuantity))
	capitalInvestedForMonthlyDesiredIncome := math.Round(calculateDesiredMonthlyIncome(asset.Price, asset.Yield, desiredMonthlyIncome))

	fmt.Printf("\n|--------------------------------------------------------------------------|")
	fmt.Printf("\n|-> Asset Name: %s", asset.Symbol)
	fmt.Printf("\n|-> Asset Price: R$%.2f", asset.Price)
	fmt.Printf("\n|-> Asset Yield: R$%.3f", asset.Yield)
	fmt.Println("\n|--------------------------------------------------------------------------|")
	fmt.Printf("|- Asset Quantity for Snowball Effect: %.0f\n", assetQuantity)
	fmt.Printf("|- Capital Invested for SnowBall Effect: R$ %.2f \n", capitalInvested)
	fmt.Println("|--------------------------------------------------------------------------|")
	fmt.Printf("|- Capital Invested for Monthly Desired Income of R$ 1000: R$ %.2f \n", capitalInvestedForMonthlyDesiredIncome)
	fmt.Println("|--------------------------------------------------------------------------|")
}

func readAssetSymbolFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the symbol of the asset: ")
	assetSymbol, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	assetSymbol = strings.TrimSpace(assetSymbol)
	assetSymbol = strings.ToUpper(assetSymbol)

	fmt.Printf("\nLoading...\n")

	return assetSymbol
}

/*calculateMagicNumber calculates the amount of stocks need to buy the stock itself with the dividends
MagicNumber: stockPrice / stockYield = stockQuantity
*/
func calculateMagicNumber(stockPrice float64, stockYield float64) float64 {
	return (stockPrice / stockYield)
}

//calculateCapitalInvestedSnowball calculates amount of money needed for the snowball effect
func calculateCapitalInvestedSnowball(stockPrice float64, stockQuantity float64) float64 {
	return (stockPrice * stockQuantity)
}

//calculateDesiredMonthlyIncome calculates the amount of capital needed to reach the desired monthly income from dividends
func calculateDesiredMonthlyIncome(stockPrice float64, stockYield float64, desiredMonthlyIncome float64) float64 {
	capitalForDesiredMonthlyIncome := (desiredMonthlyIncome / stockYield) * stockPrice

	return capitalForDesiredMonthlyIncome
}
