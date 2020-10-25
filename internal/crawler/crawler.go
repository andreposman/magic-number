package crawler

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// )

// type AssetString struct {
// 	Symbol string
// 	Price  string
// 	Yield  string
// }

// type Asset struct {
// 	Symbol string
// 	Price  float64
// 	Yield  float64
// }

// // GetAsset creates the asset struct
// func GetAsset(assetSymbol string) *AssetString {
// 	asset := new(AssetString)

// 	asset.Symbol = assetSymbol
// 	asset.Price = "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong"
// 	asset.Yield = "#main-2 > div.container.pb-7 > div:nth-child(6) > div > div > div:nth-child(1) > div > div > strong"

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
// 	element := GetAsset(assetSymbol).Yield

// 	symbolYieldString := html.Find(element).Text()
// 	symbolYieldString = strings.Replace(symbolYieldString, ",", ".", 1)

// 	symbolYield, err := strconv.ParseFloat(symbolYieldString, 5)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return symbolYield
// }

// //BuildAsset builds the asset info
// func BuildAsset(assetSymbol string) *Asset {
// 	asset := new(Asset)

// 	asset.Symbol = assetSymbol
// 	asset.Price = GetSymbolPrice(assetSymbol)
// 	asset.Yield = GetSymbolYield(assetSymbol)

// 	return asset
// }

// //BuildAsset builds the asset info
// func BuildAssetStr(assetSymbol string) *AssetString {
// 	asset := new(AssetString)

// 	asset.Symbol = assetSymbol
// 	asset.Price = ""
// 	asset.Yield = ""

// 	return asset
// }
