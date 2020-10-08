package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/andreposman/magic-number/internal/asset/investment"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/olekukonko/tablewriter"
)

type AssetTable struct {
	Symbol                      string
	Name                        string
	Price                       string
	YieldAvarage24M             string
	DividendYield               string
	MinPrice52Week              string
	MaxPrice52Week              string
	PerformanceLast12M          string
	PerformanceThisMonth        string
	DesiredMonthlyIncome        string
	MagicNumber                 string
	CapitalSnowBallEffect       string
	CapitalDesiredMonthlyIncome string
}

//PrintDataTable print a "pretty" table on the terminal
func PrintDataTable(asset *assetModel.Asset, desiredMonthlyIncome float64) {
	goals := investment.CalculateGoals(asset.Price, asset.YieldAvarage24M, desiredMonthlyIncome)
	a := convertAssetToString(asset, goals, desiredMonthlyIncome)

	dataAsset := [][]string{
		[]string{
			a.Symbol,
			a.Name,
			a.Price,
			a.YieldAvarage24M,
			a.DividendYield,
			a.MinPrice52Week,
			a.MaxPrice52Week,
			a.PerformanceLast12M,
			a.PerformanceThisMonth,
		}}

	tableAsset := tablewriter.NewWriter(os.Stdout)
	tableAsset.SetHeader([]string{
		"ASSET \nSYMBOL",
		"ASSET \nNAME",
		"ASSET \nPRICE",
		"ASSET \nYIELD AVARAGE (24M)",
		"DIVIDEND \nYIELD",
		"MIN PRICE \n52 WK",
		"MAX PRICE \n52 WK",
		"PERFORMANCE \nLAST 12M",
		"PERFORMANCE \nTHIS MONTH",
	})
	tableAsset.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	tableAsset.SetAutoWrapText(true)
	tableAsset.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	tableAsset.SetAlignment(tablewriter.ALIGN_CENTER)
	tableAsset.AppendBulk(dataAsset)
	tableAsset.Render()
	fmt.Print("\n")

	dataCapital := [][]string{
		[]string{
			a.MagicNumber,
			a.CapitalSnowBallEffect,
			a.CapitalDesiredMonthlyIncome,
		}}

	tableCapital := tablewriter.NewWriter(os.Stdout)
	tableCapital.SetHeader([]string{
		"MAGIC NUMBER\n(ASSET QTY)",
		"CAPITAL FOR\nSNOWBALL EFFECT",
		"CAPITAL FOR\nDESIRED MONTHLY INCOME OF R$" + a.DesiredMonthlyIncome,
	})
	tableCapital.SetBorders(tablewriter.Border{Left: true, Top: true, Right: true, Bottom: true})
	tableCapital.SetAutoWrapText(false)
	tableCapital.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	tableCapital.SetAlignment(tablewriter.ALIGN_CENTER)
	tableCapital.AppendBulk(dataCapital)
	tableCapital.Render()
	fmt.Print("\n")
}

func convertAssetToString(asset *assetModel.Asset, goals *investment.Goals, desiredMonthlyIncome float64) *AssetTable {
	a := new(AssetTable)

	a.Symbol = asset.Symbol
	a.Name = asset.Name
	a.Price = "R$" + strconv.FormatFloat(asset.Price, 'f', 2, 64)
	a.YieldAvarage24M = "R$" + strconv.FormatFloat(asset.YieldAvarage24M, 'f', 5, 64)
	a.DividendYield = strconv.FormatFloat(asset.DividendYield, 'f', 2, 64) + "%"
	a.MinPrice52Week = "R$" + strconv.FormatFloat(asset.MinPrice52Week, 'f', 2, 64)
	a.MaxPrice52Week = "R$" + strconv.FormatFloat(asset.MaxPrice52Week, 'f', 2, 64)
	a.PerformanceLast12M = asset.PerformanceLast12M
	a.PerformanceThisMonth = asset.PerformanceThisMonth
	a.DesiredMonthlyIncome = strconv.FormatFloat(desiredMonthlyIncome, 'f', 3, 64)
	a.MagicNumber = strconv.FormatFloat(goals.MagicNumber, 'f', 0, 64)
	a.CapitalSnowBallEffect = "R$" + strconv.FormatFloat(goals.CapitalSnowBallEffect, 'f', 2, 64)
	a.CapitalDesiredMonthlyIncome = "R$" + strconv.FormatFloat(goals.CapitalDesiredMonthlyIncome, 'f', 5, 64)

	return a
}
