package repository

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	asset "github.com/andreposman/magic-number/internal/asset/model"
	assetModel "github.com/andreposman/magic-number/internal/asset/model"
	adapter "github.com/andreposman/magic-number/internal/asset/repository/adapter"
	crawlerModel "github.com/andreposman/magic-number/internal/crawler/model"
)

//buildCrawler ...
func buildCrawler(symbol string) *crawlerModel.Crawler {
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

	asset.Symbol = req.AssetSymbol
	asset.Name = HTML.Find(assetElements.Name).Text()
	asset.Price = strings.Replace(HTML.Find(assetElements.Price).Text(), ",", ".", 1)
	asset.YieldAverage24M = strings.Replace(HTML.Find(assetElements.YieldAverage24M).Text(), ",", ".", 1)
	asset.DividendYield = strings.Replace(HTML.Find(assetElements.DividendYield).Text(), ",", ".", 1) + "%"
	asset.MinPrice52Week = strings.Replace(HTML.Find(assetElements.MinPrice52Week).Text(), ",", ".", 1)
	asset.MaxPrice52Week = strings.Replace(HTML.Find(assetElements.MaxPrice52Week).Text(), ",", ".", 1)
	asset.PerformanceLast12M = HTML.Find(assetElements.PerformanceLast12M).Text()
	asset.PerformanceThisMonth = HTML.Find(assetElements.PerformanceThisMonth).Text()
	asset.Goals.DesiredMonthlyIncome = req.DesiredMonthlyIncome

	return asset
}

//AssetCalculation returns the asset for calculation
func AssetCalculation(asset *asset.Asset) *assetModel.AssetDataFloat {
	if len(asset.Symbol) < 0 || len(asset.Price) < 0 || len(asset.YieldAverage24M) < 0 {
		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Price, Yield or Monthly Income is equal or lower to R$ 0,00")
		os.Exit(-1)
	}

	price := toExpectedFloat(asset.Price)
	yieldAverage24M := toYieldAverageFloat(asset.YieldAverage24M)
	desiredMonthlyIncome := toExpectedFloat(asset.Goals.DesiredMonthlyIncome)

	if price <= 0 && yieldAverage24M <= 0 && desiredMonthlyIncome <= 0 {
		fmt.Fprintf(os.Stderr, "\n\nError: %v\n", "Price/Yield/Desired Monthly Income must be greater than zero")
		os.Exit(-1)
	}

	a := assetModel.AssetDataFloat{
		Price:                price,
		YieldAvarage24M:      yieldAverage24M,
		DesiredMonthlyIncome: desiredMonthlyIncome,
	}

	return &a
}

//GetAsset receives asset data in string and calculations in float and returns the asset as a whole
func GetAsset(asset *asset.Asset, calculationData *assetModel.AssetDataFloat) *assetModel.Asset {

	asset.Goals.MagicNumber = strings.Replace(toExpectedString(calculationData.MagicNumber), ".00", "", 1)
	asset.Goals.CapitalSnowBallEffect = toExpectedString(calculationData.CapitalSnowBallEffect)
	asset.Goals.CapitalDesiredMonthlyIncome = toExpectedString(calculationData.CapitalDesiredMonthlyIncome)
	asset.Goals.AssetQuantityDesiredIncome = strings.Replace(toExpectedString(calculationData.AssetQuantityDesiredIncome), ".00", "", 1)

	return asset
}

// toExpectedString ...
func toExpectedString(value float64) string {
	valueString := strconv.FormatFloat(value, 'f', 2, 64)

	return valueString
}

// toYieldAvarageString ...
func toYieldAvarageString(value float64) string {
	valueString := strconv.FormatFloat(value, 'f', 5, 64)

	return valueString
}

// toMagicNumberString ...
func toMagicNumberString(value float64) string {
	valueString := strconv.FormatFloat(value, 'f', 0, 64)

	return valueString
}

// toExpectedFloat ...
func toExpectedFloat(text string) float64 {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println(err)
	}

	return value
}

// toYieldAverageFloat ...
func toYieldAverageFloat(text string) float64 {
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println(err)
	}

	return value
}
