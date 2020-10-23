package controller

import (
	asset "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/asset/service"
)

//CalculateGoals ...
func CalculateGoals(asset *asset.Model) *asset.Model {
	s := service.CalculateInvestmentGoals(asset)

	return s
}

//GetAsset ...
func GetAsset(req *asset.Request) *asset.Model {
	a := service.GetAsset(req)

	return a
}

//ReturnJSON ...
func ReturnJSON(req *asset.Model) []byte {
	a := service.BuildJSON(req)

	return a
}
