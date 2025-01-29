package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/andreposman/magic-number/internal/models"
)

func FIIHandler(ticker string) *models.FII {
	fii, err := getFII(ticker)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info(" | FII DATA | ",
		" | ID", fii.ID,
		" | Ticker", fii.Ticker,
		" | FullName", fii.FullName,
		" | Price", fii.Price,
		" | DividendYield", fii.DividendYield,
		" | Segment", fii.Segment,
	)

	return fii
}

func getFII(ticker string) (*models.FII, error) {
	fiiList, err := fetchFIIList(ticker)
	if err != nil {
		return nil, err
	}

	if len(fiiList.FullName) == 0 {
		return nil, fmt.Errorf("no FII data found")
	}

	fiiDetails, err := getFIIDetails(fiiList)
	if err != nil {
		return nil, err
	}

	return fiiDetails, nil
}

func fetchFIIList(ticker string) (*models.FII, error) {
	fiiURL := "https://investidor10.com.br/api/search/" + ticker

	res, err := http.Get(fiiURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code is %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var fiiData []*models.FII
	if err := json.Unmarshal(body, &fiiData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return fiiData[0], nil
}

func getFIIDetails(fii *models.FII) (*models.FII, error) {
	fiiDetailsURL := "https://investidor10.com.br/api/fii/comparador/table/"
	fiiID := strconv.Itoa(fii.ID)

	res, err := http.Get(fiiDetailsURL + fiiID)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code is %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var detailsResponse models.FIIDetailsResponse
	if err := json.Unmarshal(body, &detailsResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	for _, detail := range detailsResponse.Data {
		if detail.Title == fii.Ticker {
			fii.DividendYield = detail.DividendYield
			fii.Segment = detail.Segment
			fii.Type = detail.Type
			break
		}
	}

	return fii, nil
}
