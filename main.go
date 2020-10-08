package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/andreposman/magic-number/cmd"
	"github.com/andreposman/magic-number/internal/crawler/controller"
)

/*
List of TODOs:
	// TODO: struct for asset
	// TODO: modularize program (table writer, crawler, bussiness rules)
	// TODO: pattern for folders and code
	// TODO: Find out if is viable format float to BRL
	TODO: asset controller/service
	TODO: more validations
	TODO: logs
	TODO: README
*/

func main() {
	// a := config.DebugData(true)
	desiredMonthlyIncome := 1.000
	symbol := readAssetSymbolFromTerminal()
	asset := controller.GetAsset(symbol)

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
