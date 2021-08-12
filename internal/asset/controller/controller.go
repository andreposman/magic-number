package controller

import (
	asset "github.com/andreposman/magic-number/internal/asset/model"
	"github.com/andreposman/magic-number/internal/asset/service"
)

func GetAsset(req *asset.Request) (*asset.Asset, error) {
	result, err := service.GetAsset(req)
	if err != nil {
		return nil, err
	}

	return result, nil

}

//ReturnJSON ...
func ReturnJSON(req *asset.Asset) []byte {
	return service.BuildJSON(req)
}
