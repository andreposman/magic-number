package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/andreposman/magic-number/internal/config"

	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	model "github.com/andreposman/magic-number/internal/crawler/model"
)

//buildCrawler ...
func buildCrawler(req string) *model.Crawler {
	req = strings.TrimSpace(strings.ToUpper(req))

	crawler := new(model.Crawler)
	crawler.AssetSymbol = req
	crawler.URL = config.GetURL() + crawler.AssetSymbol

	return crawler
}

//GetHTML fetches the HTML document from the page
func GetHTML(req string) *goquery.Document {
	NotFoundElement := "#main-2 > section > div > h1"

	Crawler := buildCrawler(req)

	// Request the HTML page.
	res, err := http.Get(Crawler.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	HTML, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	AssetNotFound := HTML.Find(NotFoundElement).Text()
	if len(AssetNotFound) > 0 {
		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Asset does not exist")
		os.Exit(0)
	}

	return HTML
}

//GetAsset returns the asset in float
func GetAsset(req string) *assetModel.AssetString {
	asset := findAsset(req)

	return asset
}

//findAsset finds and parse the asset from the page
func findAsset(req string) *assetModel.AssetString {
	asset := new(assetModel.AssetNumber)
	HTML := GetHTML(req)
	assetElements := assetModel.AssetString{
		Symbol:               req,
		Name:                 "#main-header > div > div > div:nth-child(1) > h1 > small",
		Price:                "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong",
		YieldAvarage24M:      "#main-2 > div.container.pb-7 > div:nth-child(6) > div > div > div:nth-child(1) > div > div > strong",
		DividendYield:        "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(4) > div > div:nth-child(1) > strong",
		MinPrice52Week:       "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(2) > div > div:nth-child(1) > strong",
		MaxPrice52Week:       "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(3) > div > div:nth-child(1) > strong",
		PerformanceLast12M:   "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(5) > div > div:nth-child(1) > strong",
		PerformanceThisMonth: "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(5) > div > div.d-flex.justify-between > div > span.sub-value > b",
	}

	symbol := strings.TrimSpace(strings.ToUpper(req))
	price, err := strconv.ParseFloat(strings.Replace(HTML.Find(assetElements.Price).Text(), ",", ".", 1), 2)
	if err != nil {
		fmt.Println(err)
	}
	yieldAvarage24M, err := strconv.ParseFloat(strings.Replace(HTML.Find(assetElements.YieldAvarage24M).Text(), ",", ".", 1), 5)
	if err != nil {
		fmt.Println(err)
	}
	dividendYield, err := strconv.ParseFloat(strings.Replace(HTML.Find(assetElements.DividendYield).Text(), ",", ".", 1), 2)
	if err != nil {
		fmt.Println(err)
	}
	minPrice52Week, err := strconv.ParseFloat(strings.Replace(HTML.Find(assetElements.MinPrice52Week).Text(), ",", ".", 1), 2)
	if err != nil {
		fmt.Println(err)
	}
	maxPrice52Week, err := strconv.ParseFloat(strings.Replace(HTML.Find(assetElements.MaxPrice52Week).Text(), ",", ".", 1), 2)
	if err != nil {
		fmt.Println(err)
	}

	asset.Symbol = symbol
	asset.Name = HTML.Find(assetElements.Name).Text()
	asset.Price = price
	asset.YieldAvarage24M = yieldAvarage24M
	asset.DividendYield = dividendYield
	asset.MinPrice52Week = minPrice52Week
	asset.MaxPrice52Week = maxPrice52Week
	asset.PerformanceLast12M = HTML.Find(assetElements.PerformanceLast12M).Text()
	asset.PerformanceThisMonth = HTML.Find(assetElements.PerformanceThisMonth).Text()

	return asset
}
