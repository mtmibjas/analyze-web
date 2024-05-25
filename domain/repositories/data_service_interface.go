package repositories

import "github.com/PuerkitoBio/goquery"

type DataRepositoriesInterface interface {
	GetURLData(url string) (*goquery.Document, error)
}
