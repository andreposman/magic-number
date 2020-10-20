package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/olekukonko/tablewriter"
)

//PrintDataTable print a "pretty" table on the terminal
func PrintDataTable(asset *assetModel.ToStringConverted) {
	fmt.Printf("\n\n")

	dataAsset := [][]string{
		[]string{
			"\n" + asset.Asset.Symbol,
			"\n" + asset.Asset.Name,
			"\n" + "R$" + asset.Asset.Price,
			"\n" + "R$" + asset.Asset.YieldAvarage24M,
			"\n" + asset.Asset.DividendYield + "%",
			"\n" + "R$" + asset.Asset.MinPrice52Week,
			"\n" + "R$" + asset.Asset.MaxPrice52Week,
			"\n" + asset.Asset.PerformanceLast12M,
			"\n" + asset.Asset.PerformanceThisMonth,
		}}

	tableAsset := tablewriter.NewWriter(os.Stdout)
	tableAsset.SetHeader([]string{
		"\nASSET SYMBOL",
		"\nASSET NAME",
		"\nASSET PRICE",
		"\nYIELD \nAVARAGE(24M)",
		"\nDIVIDEND \nYIELD",
		"\nMIN. PRICE \n52 WK",
		"\nMAX. PRICE \n52 WK",
		"\nPERFORMANCE \nLAST 12M",
		"\nPERFORMANCE \nTHIS MONTH",
	})
	tableAsset.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	tableAsset.SetCenterSeparator(".")
	tableAsset.SetColumnSeparator("|")
	tableAsset.SetRowSeparator("_")
	tableAsset.SetAutoWrapText(false)
	tableAsset.SetAutoFormatHeaders(false)
	tableAsset.SetReflowDuringAutoWrap(false)
	tableAsset.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	tableAsset.SetAlignment(tablewriter.ALIGN_CENTER)
	tableAsset.AppendBulk(dataAsset)
	tableAsset.Render()
	fmt.Printf("\n\n")

	dataMagicNumber := [][]string{
		[]string{
			asset.Investment.MagicNumber,
			"R$" + asset.Investment.CapitalSnowBallEffect,
		}}

	tableMagicNumber := tablewriter.NewWriter(os.Stdout)
	tableMagicNumber.SetHeader([]string{
		"MAGIC NUMBER\n(ASSET QTY)",
		"CAPITAL FOR\nSNOWBALL EFFECT",
	})
	tableMagicNumber.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	tableMagicNumber.SetAutoWrapText(true)
	tableMagicNumber.SetCenterSeparator(".")
	tableMagicNumber.SetColumnSeparator("|")
	tableMagicNumber.SetRowSeparator("_")
	tableMagicNumber.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	tableMagicNumber.SetAlignment(tablewriter.ALIGN_CENTER)
	tableMagicNumber.AppendBulk(dataMagicNumber)
	tableMagicNumber.Render()
	fmt.Printf("\n\n")

	dataCapital := [][]string{
		[]string{
			"R$" + asset.Investment.CapitalDesiredMonthlyIncome,
		}}

	tableCapital := tablewriter.NewWriter(os.Stdout)
	tableCapital.SetHeader([]string{
		"CAPITAL FOR\nDESIRED MONTHLY INCOME OF R$" + asset.Investment.DesiredMonthlyIncome,
	})
	tableCapital.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	tableCapital.SetAutoWrapText(true)
	tableCapital.SetCenterSeparator(".")
	tableCapital.SetColumnSeparator("|")
	tableCapital.SetRowSeparator("_")
	tableCapital.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	tableCapital.SetAlignment(tablewriter.ALIGN_CENTER)
	tableCapital.AppendBulk(dataCapital)
	tableCapital.Render()
	fmt.Printf("\n\n")
}

//ReadAssetSymbolFromTerminal ...
func ReadAssetSymbolFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> Enter the symbol of the asset: ")
	assetSymbol, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	assetSymbol = strings.TrimSpace(assetSymbol)
	assetSymbol = strings.ToUpper(assetSymbol)

	return assetSymbol
}

//ReadDesiredMonthlyIncomeFromTerminal ...
func ReadDesiredMonthlyIncomeFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n-> Enter the desired monthly income: ")
	desiredMonthlyIncome, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	desiredMonthlyIncome = strings.TrimSpace(desiredMonthlyIncome)
	desiredMonthlyIncome = strings.ToUpper(desiredMonthlyIncome)

	f, err := strconv.ParseFloat(desiredMonthlyIncome, 32)
	if f <= 0 || err != nil {
		fmt.Printf("\nDesired monthly income must be a number, greater than 0.\n\n")
		os.Exit(-1)
	}

	fmt.Printf("\nLoading...\n")

	return desiredMonthlyIncome
}
