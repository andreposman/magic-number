package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andreposman/magic-number/cmd"
	asset "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/crawler/controller"
)

/*
List of TODOs:
	// TODO: struct for asset
	// TODO: modularize program (table writer, crawler, bussiness rules)
	// TODO: pattern for folders and code
	// TODO: Find out if is viable format float to BRL
	TODO: fix types and conversion on asset service
	TODO: asset controller/service
	TODO: more validations
	TODO: logs
	TODO: README
*/

func main() {
	// a := config.DebugData(true)
	// desiredMonthlyIncome := 1.000
	// desiredMonthlyIncome := readAssetSymbolFromTerminal()
	// symbol := readAssetSymbolFromTerminal()

	request := new(asset.RequestAsset)
	request.AssetSymbol = readAssetSymbolFromTerminal()
	request.DesiredMonthlyIncome = readAssetSymbolFromTerminal()

	asset := controller.GetAsset(req)

	cmd.PrintDataTable(asset, desiredMonthlyIncome)
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

func readDesiredMonthlyIncomeFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the desired monthly income: ")
	desiredMonthlyIncome, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	desiredMonthlyIncome = strings.TrimSpace(desiredMonthlyIncome)
	desiredMonthlyIncome = strings.ToUpper(desiredMonthlyIncome)

	fmt.Printf("\nLoading...\n\n\n")

	return desiredMonthlyIncome
}
