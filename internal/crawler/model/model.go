package modelcrawler

import "github.com/PuerkitoBio/goquery"

type Crawler struct {
	URL         string
	AssetSymbol string
}

type Result struct {
	HTML *goquery.Document
}
