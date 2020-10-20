package main

import (
	"github.com/andreposman/magic-number/cmd"
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
	TODO: Fix names
	TODO: logs
	TODO: README
*/

func main() {
	request := new(asset.Request)
	// request.AssetSymbol = config.DebugData().Asset
	// request.DesiredMonthlyIncome = config.DebugData().DesiredMonthlyIncome

	request.AssetSymbol = cmd.ReadAssetSymbolFromTerminal()
	request.DesiredMonthlyIncome = cmd.ReadDesiredMonthlyIncomeFromTerminal()

	asset := controller.GetAsset(request)

	cmd.PrintDataTable(asset)
	controller.ReturnJSON(asset)
}
