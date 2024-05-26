package repositories

import (
	"analyze-web/pkg/logger/zap"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (dr *DataRepository) GetURLData(url string) (*goquery.Document, error) {

	resp, err := dr.HTTPClient.Get(url)
	if err != nil {
		zap.Error("repo:GetURLData:", err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to fetch URL: %s", resp.Status)
		zap.Error("repo:GetURLData:", err)
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		zap.Error("repo:GetURLData:", err)
		return nil, err
	}

	return doc, nil
}
