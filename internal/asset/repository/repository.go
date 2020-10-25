package repository

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	adapter "github.com/andreposman/magic-number/internal/asset/repository/adapter"
	crawlerModel "github.com/andreposman/magic-number/internal/crawler/model"
)

// asset.service.func -> asset.repository.func -> crawler.service.func

//buildCrawler ...
func buildCrawler(symbol string) *crawlerModel.Crawler {
	// crawler := new(crawlerModel.Crawler)
	symbol = strings.TrimSpace(strings.ToUpper(symbol))
	crawler := crawlerModel.Crawler{
		AssetSymbol: symbol,
		URL:         adapter.StatusInvestData(symbol).URL,
	}

	return &crawler
}

//ScrapeHTML fetches the HTML document from the page
func ScrapeHTML(symbol string) *goquery.Document {
	NotFoundElement := adapter.StatusInvestData(symbol).AssetNotFound
	Crawler := buildCrawler(symbol)

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
		os.Exit(-1)
	}

	return HTML
}

//FindAsset finds and parse the asset from the page
func FindAsset(req *assetModel.Request) *assetModel.Asset {
	asset := new(assetModel.Asset)
	symbol := req.AssetSymbol
	HTML := ScrapeHTML(symbol)

	assetElements := adapter.StatusInvestElements(symbol)
	// assetElements := assetModel.Elements{
	// 	Symbol:               req.AssetSymbol,
	// 	Name:                 "#main-header > div > div > div:nth-child(1) > h1 > small",
	// 	Price:                "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong",
	// 	YieldAverage24M:      "#main-2 > div.container.pb-7 > div:nth-child(6) > div > div > div:nth-child(1) > div > div > strong",
	// 	DividendYield:        "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(4) > div > div:nth-child(1) > strong",
	// 	MinPrice52Week:       "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(2) > div > div:nth-child(1) > strong",
	// 	MaxPrice52Week:       "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(3) > div > div:nth-child(1) > strong",
	// 	PerformanceLast12M:   "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(5) > div > div:nth-child(1) > strong",
	// 	PerformanceThisMonth: "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(5) > div > div.d-flex.justify-between > div > span.sub-value > b",
	// }

	asset.Symbol = req.AssetSymbol
	asset.Name = HTML.Find(assetElements.Name).Text()
	asset.Price = strings.Replace(HTML.Find(assetElements.Price).Text(), ",", ".", 1)
	asset.YieldAverage24M = strings.Replace(HTML.Find(assetElements.YieldAverage24M).Text(), ",", ".", 1)
	asset.DividendYield = strings.Replace(HTML.Find(assetElements.DividendYield).Text(), ",", ".", 1)
	asset.MinPrice52Week = strings.Replace(HTML.Find(assetElements.MinPrice52Week).Text(), ",", ".", 1)
	asset.MaxPrice52Week = strings.Replace(HTML.Find(assetElements.MaxPrice52Week).Text(), ",", ".", 1)
	asset.PerformanceLast12M = HTML.Find(assetElements.PerformanceLast12M).Text()
	asset.PerformanceThisMonth = HTML.Find(assetElements.PerformanceThisMonth).Text()
	asset.Goals.DesiredMonthlyIncome = req.DesiredMonthlyIncome

	return asset
}

// // GetAsset creates the asset struct
// func GetAsset(assetSymbol string) *assetModel.Asset {
// 	asset := new(assetModel.Asset)

// 	asset.Symbol = assetSymbol
// 	asset.Price = "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong"
// 	asset.YieldAverage24M = "#main-2 > div.container.pb-7 > div:nth-child(6) > div > div > div:nth-child(1) > div > div > strong"

// 	return asset
// }

// //GetWebsiteHTML fetches the HTML document from the page
// func GetWebsiteHTML(assetSymbol string) *goquery.Document {

// 	asset := GetAsset(assetSymbol)

// 	webSiteURL := "https://statusinvest.com.br/fundos-imobiliarios/" + asset.Symbol
// 	pageNotFoundElement := "#main-2 > section > div > h1"

// 	// Request the HTML page.
// 	res, err := http.Get(webSiteURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	if res.StatusCode != 200 {
// 		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
// 	}

// 	// Load the HTML document
// 	html, err := goquery.NewDocumentFromReader(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pageNotFound := html.Find(pageNotFoundElement).Text()

// 	if len(pageNotFound) > 0 {
// 		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Asset does not exist")
// 		os.Exit(0)
// 	}

// 	return html
// }

// //GetSymbolPrice fetch the price of the symbol in Brazilian Reais
// func GetSymbolPrice(assetSymbol string) float64 {
// 	html := GetWebsiteHTML(assetSymbol)
// 	element := GetAsset(assetSymbol).Price

// 	symbolPriceString := html.Find(element).Text()
// 	symbolPriceString = strings.Replace(symbolPriceString, ",", ".", 1)

// 	symbolPrice, err := strconv.ParseFloat(symbolPriceString, 3)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return symbolPrice
// }

// //GetSymbolYield fetch the yield of the symbol in Brazilian Reais
// func GetSymbolYield(assetSymbol string) float64 {
// 	html := GetWebsiteHTML(assetSymbol)
// 	element := GetAsset(assetSymbol).YieldAverage24M

// 	symbolYieldString := html.Find(element).Text()
// 	symbolYieldString = strings.Replace(symbolYieldString, ",", ".", 1)

// 	symbolYield, err := strconv.ParseFloat(symbolYieldString, 5)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return symbolYield
// }
