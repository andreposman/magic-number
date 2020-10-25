package adapter

import assetModel "github.com/andreposman/magic-number/internal/asset/model"

type StatusInvest struct {
	URL           string
	AssetNotFound string
}

// StatusInvestData returns url and not found error
func StatusInvestData(symbol string) *StatusInvest {
	statusInvest := StatusInvest{
		URL:           "https://statusinvest.com.br/fundos-imobiliarios/" + symbol,
		AssetNotFound: "#main-2 > section > div > h1",
	}

	return &statusInvest
}

//StatusInvestElements returns the html selectors for the crawler to scrape it
func StatusInvestElements(symbol string) *assetModel.Elements {
	asset := assetModel.Elements{
		Symbol:               symbol,
		Name:                 "#main-header > div > div > div:nth-child(1) > h1 > small",
		Price:                "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div.info.special.w-100.w-md-33.w-lg-20 > div > div:nth-child(1) > strong",
		YieldAverage24M:      "#main-2 > div.container.pb-7 > div:nth-child(6) > div > div > div:nth-child(1) > div > div > strong",
		DividendYield:        "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(4) > div > div:nth-child(1) > strong",
		MinPrice52Week:       "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(2) > div > div:nth-child(1) > strong",
		MaxPrice52Week:       "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(3) > div > div:nth-child(1) > strong",
		PerformanceLast12M:   "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(5) > div > div:nth-child(1) > strong",
		PerformanceThisMonth: "#main-2 > div.container.pb-7 > div.top-info.d-flex.flex-wrap.justify-between.mb-3.mb-md-5 > div:nth-child(5) > div > div.d-flex.justify-between > div > span.sub-value > b",
	}

	return &asset
}
