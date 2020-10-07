package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andreposman/magic-number/internal/config"
	"github.com/andreposman/magic-number/internal/crawler/controller"
)

/*
List of TODOs:
	TODO: struct for asset
	TODO: logs
	TODO: modularize program (table writer, crawler, bussiness rules)
	TODO: pattern for folders and code
	TODO: more validations
	TODO: README
	TODO: Find out if is viable format float to BRL
*/

func main() {
	a := config.DebugData(true)
	controller.GetAsset(a)

	// assetSymbol := readAssetSymbolFromTerminal()
	// // assetSymbol := "HGLG11"

	// var desiredMonthlyIncome float64 = 1000
	// asset := crawler.BuildAsset(assetSymbol)

	// assetQuantityMagicNumber := math.Round(calculateMagicNumber(asset.Price, asset.YieldAvarage24M))
	// capitalInvestedSnowBall := (calculateCapitalInvestedSnowball(asset.Price, assetQuantityMagicNumber))
	// capitalInvestedForMonthlyDesiredIncome := (calculateDesiredMonthlyIncome(asset.Price, asset.YieldAvarage24M, desiredMonthlyIncome))

	// assetStr := crawler.BuildAssetStr(assetSymbol)
	// desiredMonthlyIncomeStr := strconv.FormatFloat(desiredMonthlyIncome, 'f', 2, 64)
	// assetStr.Symbol = asset.Symbol
	// assetStr.Price = "R$ " + strconv.FormatFloat(asset.Price, 'f', 3, 64)
	// assetStr.Yield = "R$ " + strconv.FormatFloat(asset.Yield, 'f', 5, 64)

	// assetQuantityMagicNumberStr := strconv.FormatFloat(assetQuantityMagicNumber, 'f', 0, 64)
	// capitalInvestedSnowBallStr := "R$ " + strconv.FormatFloat(capitalInvestedSnowBall, 'f', 2, 64)
	// capitalInvestedForMonthlyDesiredIncomeStr := "R$ " + strconv.FormatFloat(capitalInvestedForMonthlyDesiredIncome, 'f', 2, 64)

	// data := [][]string{
	// 	[]string{
	// 		// assetStr.Symbol,
	// 		"R$ " + strconv.FormatFloat(asset.Price, 'f', 3, 64),
	// 		// asset.YieldAvarage24M,
	// 		// assetQuantityMagicNumber,
	// 		// capitalInvestedSnowBall,
	// 		// capitalInvestedForMonthlyDesiredIncome
	// 	}}

	// table := tablewriter.NewWriter(os.Stdout)
	// table.SetHeader([]string{
	// 	"ASSET\nNAME",
	// 	"ASSET\nPRICE",
	// 	"ASSET YIELD\n(24M)",
	// 	"MAGIC NUMBER\n(QTY)",
	// 	"CAPITAL FOR\nSNOWBALL EFFECT",
	// "CAPITAL FOR \n DESIRED MONTHLY INCOME OF R$" + desiredMonthlyIncome
	// })
	// table.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	// table.SetAutoWrapText(false)
	// table.SetCenterSeparator("|")
	// table.SetRowSeparator("-")
	// table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	// table.SetAlignment(tablewriter.ALIGN_CENTER)
	// table.AppendBulk(data)
	// table.Render()
	// fmt.Print("\n")
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

	fmt.Printf("\nLoading...\n\n\n")

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
