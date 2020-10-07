package controller

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	service "github.com/andreposman/magic-number/internal/crawler/service"
)

//GetHTML receives req and returns the html
func GetHTML(req string) *goquery.Document {
	res := service.GetHTML(req)

	return res
}

func GetAsset(req string) *assetModel.Asset {
	res := service.GetAsset(req)

	fmt.Println(res)
	return res
}
