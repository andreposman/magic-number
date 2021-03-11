package main

import (
	"github.com/andreposman/magic-number/delivery/api"
	"github.com/andreposman/magic-number/delivery/terminal"
	"github.com/andreposman/magic-number/internal/asset/controller"
	asset "github.com/andreposman/magic-number/internal/asset/model"
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
	// TODO: Fix Bug when desired income is 1
	// TODO: API
	// TODO: Review code (func names and code organization)
	TODO: Improve Tests
	TODO: replace possible commas for dots in desired income
	TODO: Logs
	TODO: Docker
	TODO: README
	TODO: start the Flutter app and consume the api
*/

func main() {

	request := new(asset.Request)
	// request.AssetSymbol = config.DebugData().Asset
	// request.DesiredMonthlyIncome = config.DebugData().DesiredMonthlyIncome

	request.AssetSymbol = terminal.ReadAssetSymbolFromTerminal()
	request.DesiredMonthlyIncome = terminal.ReadDesiredMonthlyIncomeFromTerminal()

	asset := controller.GetAsset(request)

	terminal.PrintDataTable(asset)
	api.Init(asset)
}
