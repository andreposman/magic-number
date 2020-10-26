package controller

import (
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/asset/service"
)

//GetAsset ...
func GetAsset(req *assetModel.Request) *assetModel.Asset {
	return service.GetAsset(req)
}

//ReturnJSON ...
func ReturnJSON(req *assetModel.Asset) []byte {
	return service.BuildJSON(req)
}
