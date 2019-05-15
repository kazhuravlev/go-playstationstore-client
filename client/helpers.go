package client

import (
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"strings"
)

type SrcSet map[float64]string

func (s SrcSet) GetMax() string {
	var max float64
	for k := range s {
		if k > max {
			max = k
		}
	}

	return s[max]
}

func getPageNum(s *goquery.Selection) (int, error) {
	href, _ := s.Attr("href")
	values := strings.Split(href, "/")
	pageNum, err := strconv.Atoi(values[len(values)-1])
	if err != nil {
		return 0, err
	}

	return pageNum, nil
}

func getSrcSet(srcSet string) (SrcSet, error) {
	res := make(SrcSet)
	for _, spec := range strings.Split(srcSet, ",") {
		spec = strings.Trim(spec, " ")
		values := strings.SplitAfterN(spec, " ", 2)
		scale, err := strconv.ParseFloat(strings.Trim(values[1], "x@"), 64)
		if err != nil {
			return nil, err
		}

		res[scale] = values[0]
	}

	return res, nil
}
