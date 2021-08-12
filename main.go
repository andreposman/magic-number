package main

import (
	"github.com/andreposman/magic-number/delivery/api"
	"github.com/andreposman/magic-number/delivery/terminal"
	"github.com/andreposman/magic-number/internal/asset/controller"
	asset "github.com/andreposman/magic-number/internal/asset/model"
)

/*
List of TODOs:
	TODO: Improve Tests
	TODO: replace possible commas for dots in desired income
	TODO: Logs
	TODO: Docker
	TODO: README
	TODO: start the Flutter app and consume the api
*/

func main() {

	var request asset.Request
	// request.AssetSymbol = config.DebugData().Asset
	// request.DesiredMonthlyIncome = config.DebugData().DesiredMonthlyIncome

	request.AssetSymbol = terminal.ReadAssetSymbolFromTerminal()
	request.DesiredMonthlyIncome = terminal.ReadDesiredMonthlyIncomeFromTerminal()

	asset, err := controller.GetAsset(&request)
	if err != nil {
		panic(err)
	}

	terminal.PrintDataTable(asset)
	api.Init(asset)
}
