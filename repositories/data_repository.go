package repositories

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func (dr *DataRepository) GetURLData(url string) (*goquery.Document, error) {

	resp, err := dr.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	// content, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, err
	// }
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch URL: %s", resp.Status)
	}
	// if _, err := html.Parse(resp.Body); err != nil {
	// 	fmt.Println("sfdsfs")
	// }
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
