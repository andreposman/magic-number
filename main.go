package main

import (
	"github.com/andreposman/magic-number/delivery/terminal"
	"github.com/andreposman/magic-number/internal/asset/controller"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
)

/*
List of TODOs:
	// TODO: struct for asset
	// TODO: modularize program (table writer, crawler, bussiness rules)
	// TODO: pattern for folders and code
	// TODO: Find out if is viable format float to BRL
	// TODO: fix types and conversion on asset service
	// TODO: asset controller/service
	// TODO: more validations
	// TODO: return in json
	// TODO: implement repository pattern
	// TODO: Fix names
	TODO: Review code (func names and code organization)
	TODO: Tests
	TODO: replace possible commas for dots in desired income
	TODO: Fix Bug when desired income is 1
	TODO: README
	TODO: API
	TODO: Logs
	TODO: Docker
*/

func main() {
	request := new(assetModel.Request)
	// request.AssetSymbol = config.DebugData().Asset
	// request.DesiredMonthlyIncome = config.DebugData().DesiredMonthlyIncome

	request.AssetSymbol = terminal.ReadAssetSymbolFromTerminal()
	request.DesiredMonthlyIncome = terminal.ReadDesiredMonthlyIncomeFromTerminal()

	asset := controller.GetAsset(request)

	terminal.PrintDataTable(asset)
	controller.ReturnJSON(asset)
}
