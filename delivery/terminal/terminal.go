package terminal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/olekukonko/tablewriter"
)

//PrintDataTable print a "pretty" table on the terminal
func PrintDataTable(asset *assetModel.Asset) {
	fmt.Printf("\n\n")

	dataAsset := [][]string{
		[]string{
			"\n" + asset.Symbol,
			"\n" + asset.Name,
			"\n" + "R$" + asset.Price,
			"\n" + "R$" + asset.YieldAverage24M,
			"\n" + asset.DividendYield,
			"\n" + "R$" + asset.MinPrice52Week,
			"\n" + "R$" + asset.MaxPrice52Week,
			"\n" + asset.PerformanceLast12M,
			"\n" + asset.PerformanceThisMonth,
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
			asset.Goals.MagicNumber,
			"R$" + asset.Goals.CapitalSnowBallEffect,
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
			asset.Goals.AssetQuantityDesiredIncome,
			"R$" + asset.Goals.CapitalDesiredMonthlyIncome,
		}}

	tableCapital := tablewriter.NewWriter(os.Stdout)
	tableCapital.SetHeader([]string{
		"ASSET QUANTITY \nFOR DESIRED MONTLY INCOME",
		"CAPITAL FOR\nDESIRED MONTHLY INCOME OF R$" + asset.Goals.DesiredMonthlyIncome,
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

	if strings.Contains(desiredMonthlyIncome, ",") {
		desiredMonthlyIncome = strings.Replace(desiredMonthlyIncome, ",", ".", 1)
	}

	f, err := strconv.ParseFloat(desiredMonthlyIncome, 32)
	if f <= 0 || err != nil {
		// fmt.Printf("\nDesired monthly income must be a number, greater than 0.\n\n")
		log.Fatal("\nDesired monthly income must be a number, greater than 0.\n\n")
		os.Exit(-1)
	}

	fmt.Printf("\nLoading...\n")

	return desiredMonthlyIncome
}
