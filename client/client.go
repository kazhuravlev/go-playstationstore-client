package client

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	http *http.Client
}

func New(options ...Option) (*Client, error) {
	c := Client{
		http: http.DefaultClient,
	}

	for i := range options {
		if err := options[i](&c); err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	return &c, nil
}

func (c *Client) GetAllGames(page int) ([]GameOverview, *ListingMeta, error) {
	u := fmt.Sprintf("https://store.playstation.com/ru-ru/grid/STORE-MSF75508-FULLGAMES/%d", page)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	var games []GameOverview
	doc.Find(".grid-cell--game").Each(func(i int, s *goquery.Selection) {
		game := GameOverview{}
		game.Link, _ = s.Find("a.internal-app-link").First().Attr("href")
		game.Name, _ = s.Find("div.grid-cell__title span").First().Attr("title")
		game.Price = s.Find("h3.price-display__price").First().Text()
		game.PricePsPlus = s.Find("div.price-display__price--is-plus-upsell").First().Text()

		game.ImagesSrcSet = make(SrcSet)

		srcSetStr, _ := s.Find("img.product-image__img.product-image__img--product.product-image__img-main").First().Attr("srcset")
		srcSet, err := getSrcSet(srcSetStr)
		if err != nil {
			return
		}
		game.ImagesSrcSet = srcSet

		games = append(games, game)
	})

	meta := ListingMeta{
		CurrentPage:   page,
		ObjectsOnPage: len(games),
	}

	doc.Find("div.paginator-control__container a").Each(func(i int, s *goquery.Selection) {
		if s.HasClass("paginator-control__beginning") {
			pageNum, err := getPageNum(s)
			if err != nil {
				return
			}

			meta.MinPage = pageNum
		}

		if s.HasClass("paginator-control__end") {
			pageNum, err := getPageNum(s)
			if err != nil {
				return
			}

			meta.MaxPage = pageNum
		}
	})

	return games, &meta, nil
}
