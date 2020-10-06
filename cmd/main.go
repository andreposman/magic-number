package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/andreposman/magic-number/internal/crawler"
	"github.com/olekukonko/tablewriter"
)

func main() {

	// assetSymbol := readAssetSymbolFromTerminal()
	assetSymbol := "HGLG11"

	var desiredMonthlyIncome float64 = 1000
	asset := crawler.BuildAsset(assetSymbol)

	assetQuantityMagicNumber := math.Round(calculateMagicNumber(asset.Price, asset.Yield))
	capitalInvestedSnowBall := (calculateCapitalInvestedSnowball(asset.Price, assetQuantityMagicNumber))
	capitalInvestedForMonthlyDesiredIncome := (calculateDesiredMonthlyIncome(asset.Price, asset.Yield, desiredMonthlyIncome))

	// assetString, assetQuantityMagicNumberString,  := crawler.FormatDataForDisplay(asset, assetQuantityMagicNumber, capitalInvestedSnowBall, capitalInvestedForMonthlyDesiredIncome)
	assetStr := crawler.BuildAssetStr(assetSymbol)

	assetStr.Symbol = asset.Symbol
	assetStr.Price = "R$ " + strconv.FormatFloat(asset.Price, 'f', 3, 64)
	assetStr.Yield = "R$ " + strconv.FormatFloat(asset.Yield, 'f', 5, 64)

	assetQuantityMagicNumberStr := strconv.FormatFloat(assetQuantityMagicNumber, 'f', 0, 64)
	capitalInvestedSnowBallStr := "R$ " + strconv.FormatFloat(capitalInvestedSnowBall, 'f', 2, 64)
	capitalInvestedForMonthlyDesiredIncomeStr := "R$ " + strconv.FormatFloat(capitalInvestedForMonthlyDesiredIncome, 'f', 2, 64)

	data := [][]string{
		[]string{
			assetStr.Symbol,
			assetStr.Price,
			assetStr.Yield,
			assetQuantityMagicNumberStr,
			capitalInvestedSnowBallStr,
			capitalInvestedForMonthlyDesiredIncomeStr}}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"ASSET NAME",
		"ASSET PRICE",
		"ASSET YIELD (24M)",
		"MAGIC NUMBER (QTY)",
		"CAPITAL FOR SNOWBALL EFFECT",
		"CAPITAL FOR DESIRED MONTHLY INCOME"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()

	fmt.Printf("\n|--------------------------------------------------------------------------|")
	fmt.Printf("\n|-> Asset Name: %s", asset.Symbol)
	fmt.Printf("\n|-> Asset Price: R$%.2f", asset.Price)
	fmt.Printf("\n|-> Asset Yield: R$%.3f", asset.Yield)
	fmt.Println("\n|--------------------------------------------------------------------------|")
	fmt.Printf("|- Asset Quantity for Snowball Effect: %.0f\n", assetQuantityMagicNumber)
	fmt.Printf("|- Capital Invested for SnowBall Effect: R$ %.2f \n", capitalInvestedSnowBall)
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

	r := stockPrice / stockYield

	return r
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
