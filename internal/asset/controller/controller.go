package controller

import (
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/asset/service"
)

//CalculateGoals ...
func CalculateGoals(asset *assetModel.Asset) *assetModel.Asset {
	s := service.CalculateInvestmentGoals(asset)

	return s
}

//GetAsset ...
func GetAsset(req *assetModel.Request) *assetModel.Asset {
	a := service.GetAsset(req)

	return a
}

//ReturnJSON ...
func ReturnJSON(req *assetModel.Asset) []byte {
	a := service.BuildJSON(req)

	return a
}
