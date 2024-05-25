package usecases

import (
	"analyze-web/domain/entities"
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func (s *Service) GetURLData(urlStr string) (*entities.URLData, error) {

	res, err := s.DataRepository.GetURLData(urlStr)
	if err != nil {
		return nil, err
	}
	result := &entities.URLData{
		HTMLVersion:       s.getHTMLVersion(res),
		Title:             res.Find("title").Text(),
		Headings:          s.getHeadings(res),
		InternalLinks:     0,
		ExternalLinks:     0,
		InaccessibleLinks: 0,
		ContainsLoginForm: s.containsLoginForm(res),
	}

	// _, err := url.Parse(urlStr)
	// if err != nil {
	// 	return nil, err
	// }

	res.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if !exists {
			return
		}

		linkURL, err := url.Parse(link)
		if err != nil || linkURL.IsAbs() {
			result.ExternalLinks++
		} else {
			result.InternalLinks++
		}

		if err != nil || !exists {
			result.InaccessibleLinks++
		}
	})

	return result, nil
}

func (s *Service) getHTMLVersion(res *goquery.Document) string {
	for _, n := range res.Nodes {
		if n.Type == html.DoctypeNode {
			if strings.Contains(n.Data, "html") {
				return "HTML5"
			}
			break
		}
	}
	return "UNKNOWN"
}

func (s *Service) getHeadings(res *goquery.Document) map[string]int {
	head := make(map[string]int)
	for i := 1; i <= 6; i++ {
		h := fmt.Sprintf("h%d", i)
		head[h] = res.Find(h).Length()
	}
	return head
}
func (s *Service) containsLoginForm(res *goquery.Document) bool {
	var form bool
	res.Find("form").Each(func(i int, s *goquery.Selection) {
		if s.Find("input[type='password']").Length() > 0 {
			form = true
		}
	})
	return form
}
