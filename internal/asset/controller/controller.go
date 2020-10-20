package controller

import (
	asset "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/asset/service"
)

//CalculateGoals ...
func CalculateGoals(asset *asset.Model) *asset.ToStringConverted {
	s := service.CalculateGoals(asset)

	return s
}

//GetAsset ...
func GetAsset(req *asset.Request) *asset.ToStringConverted {
	a := service.GetAsset(req)

	return a
}

//ReturnJSON ...
func ReturnJSON(req *asset.ToStringConverted) []byte {
	a := service.BuildJSON(req)

	return a
}
