package controller

import (
	asset "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/asset/service"
)

func CalculateGoals(req string, desiredMonthlyIncome float64) *asset.Goals {
	s := service.CalculateGoals(req, desiredMonthlyIncome)

	return s
}

func ConvertAssetToString(asset *asset.AssetNumber) asset.AssetString {
	// res := service.ConvertAssetToString(asset)

	// return res
}
