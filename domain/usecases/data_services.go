package usecases

import (
	"analyze-web/domain"
	"analyze-web/domain/entities"
	"analyze-web/pkg/logger/zap"
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func (s *Service) GetURLData(urlStr string) (*entities.URLData, error) {
	res, err := s.DataRepository.GetURLData(urlStr)
	if err != nil {
		zap.Error("usecase:GetURLData:", err)
		return nil, err
	}
	htmlChan := s.getHTMLVersion(res)
	titleChan := s.FindTitle(res)
	headingChan := s.getHeadings(res)
	linksChan := s.processLinks(res)
	loginFormChan := s.containsLoginForm(res)

	hTMLVersion := <-htmlChan
	title := <-titleChan
	headings := <-headingChan
	links := <-linksChan
	loginForm := <-loginFormChan
	zap.Info("usecase:GetURLData:", "sucessfully fetch data!")
	return &entities.URLData{
		HTMLVersion:       hTMLVersion,
		Title:             title,
		Headings:          headings,
		InternalLinks:     links.Internal,
		ExternalLinks:     links.External,
		InaccessibleLinks: links.Inaccessible,
		ContainsLoginForm: loginForm,
	}, nil
}

func (s *Service) getHTMLVersion(res *goquery.Document) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)

		for _, n := range res.Nodes {
			if n.FirstChild != nil && n.FirstChild.Type == html.DoctypeNode {
				docData := n.FirstChild.Data
				if docData == "html" {
					if len(n.FirstChild.Attr) > 0 {
						doctype := n.FirstChild.Attr[0].Val
						switch {
						case strings.Contains(doctype, "HTML 4"):
							ch <- "HTML 4"
						case strings.Contains(doctype, "HTML 3"):
							ch <- "HTML 3"
						default:
							ch <- "unknown"
						}
					} else {
						ch <- "HTML5"
					}
					return
				}
				if docData == "public" {
					ch <- "XHTML 1"
					return
				}
			}
		}

		ch <- "unknown"
	}()
	return ch
}
func (s *Service) FindTitle(res *goquery.Document) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		ch <- res.Find("title").Text()
	}()
	return ch
}
func (s *Service) getHeadings(res *goquery.Document) <-chan map[string]int {
	ch := make(chan map[string]int)
	go func() {
		defer close(ch)

		head := make(map[string]int)
		for i := 1; i <= 6; i++ {
			h := fmt.Sprintf("h%d", i)
			head[h] = res.Find(h).Length()
		}
		ch <- head
	}()
	return ch
}
func (s *Service) containsLoginForm(res *goquery.Document) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)

		res.Find("form").Each(func(i int, s *goquery.Selection) {
			if s.Find("input[type='password']").Length() > 0 {
				ch <- true
				return
			}
		})
		ch <- false
	}()
	return ch
}

func (s *Service) processLinks(res *goquery.Document) <-chan domain.LinkChan {
	ch := make(chan domain.LinkChan)
	go func() {
		defer close(ch)

		internalLinks, externalLinks, inaccessibleLinks := 0, 0, 0

		res.Find("a").Each(func(i int, s *goquery.Selection) {
			link, exists := s.Attr("href")
			if !exists {
				inaccessibleLinks++
				return
			}

			linkURL, err := url.ParseRequestURI(link)
			if err != nil {
				inaccessibleLinks++
				return
			}

			if linkURL.IsAbs() {
				externalLinks++
			} else {
				internalLinks++
			}
		})
		ch <- domain.LinkChan{
			Internal:     internalLinks,
			External:     externalLinks,
			Inaccessible: inaccessibleLinks,
		}
	}()
	return ch
}
