package controller

import (
	"github.com/andreposman/magic-number/internal/models"
	"github.com/andreposman/magic-number/internal/service"
)

func FetchFIIData(ticker string) (*models.FII, error) {
	fii, err := service.GetFII(ticker)
	if err != nil {
		return nil, err
	}

	return fii, nil
}
