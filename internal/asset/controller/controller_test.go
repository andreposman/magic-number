package controller_test

import (
	"testing"

	"github.com/andreposman/magic-number/internal/asset/controller"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/kinbiko/jsonassert"
)

//TestGetAsset sends an asset and expects a return
func TestGetAsset(t *testing.T) {
	req := assetModel.Request{
		AssetSymbol:          "HGLG11",
		DesiredMonthlyIncome: "1000",
	}

	result, err := controller.GetAsset(&req)
	if err != nil {
		t.Errorf("Error: ", err)

	}

	if result.Symbol != "HGLG11" {
		t.Errorf("Error: Asset does not exist")
	}
}

//TestReturnJSON sends a mocked asset and expects this return
func TestReturnJSON(t *testing.T) {
	ja := jsonassert.New(t)
	mockGoals := assetModel.Goals{
		MagicNumber:                 "231",
		CapitalSnowBallEffect:       "41187.30",
		DesiredMonthlyIncome:        "1000",
		AssetQuantityDesiredIncome:  "1297.02",
		CapitalDesiredMonthlyIncome: "231258.11",
	}
	mockAsset := assetModel.Asset{
		Symbol:               "HGLG11",
		Name:                 "CGHG Logística",
		Price:                "178.30",
		YieldAverage24M:      "0.77100000",
		DividendYield:        "6.19",
		MinPrice52Week:       "117.05",
		MaxPrice52Week:       "196.50",
		PerformanceLast12M:   "-4,27",
		PerformanceThisMonth: "0,63",
		Goals:                mockGoals,
	}

	expectedJSON :=
		`
    {
		"symbol": "HGLG11",
		"name": "CGHG Logística",
		"price": "178.30",
		"yieldAverage24m": "0.77100000",
		"dividendYield": "6.19",
		"minPrice52wk": "117.05",
		"maxPrice52wk": "196.50",
		"performanceLast12m": "-4,27",
		"performanceThisMonth": "0,63",
		"goals": {
			"magicNumber": "231",
			"capitalSnowBallEffect": "41187.30",
			"desiredMonthlyIncome": "1000",
			"assetQuantityDesiredIncome": "1297.02",
			"capitalDesiredMonthlyIncome": "231258.11"
		}
	}
	`

	result := controller.ReturnJSON(&mockAsset)

	if len(result) < 1 {
		t.Errorf("Error: Asset does not exist")
	}

	ja.Assertf(string(result), expectedJSON)
}
