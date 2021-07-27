package back

import (
	"log"
	"sort"
	"time"
)

const (
	EX_URL = "https://docs.google.com/spreadsheets/d/1ud6IZFjoT0Hyvh6TjbGq7czDezs9M3ODqEKysQ04h8E/pub?gid=1&single=true&output=csv"
	Layout = "02/01/2006"
)

type ByDate [][]string

func (s ByDate) Len() int {
	return len(s)
}
func (s ByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByDate) Less(i, j int) bool {
	a, _ := time.Parse(Layout, s[i][1])
	b, _ := time.Parse(Layout, s[j][1])
	return a.Before(b)
}

func allExpos() ([][]string, error) {
	if data, err := readCSVFromUrl(EX_URL); err != nil {
		log.Println(err)
		return nil, err
	} else {
		sort.Sort(ByDate(data))
		return data, nil
	}

	return nil, ErrKeyNotFound
}

func currentExpos() ([][]string, error) {
	data, err := allExpos()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	expos := [][]string{}

	for _, expo := range data {
		expoDate, _ := time.Parse(Layout, expo[1])
		if expoDate.After(now) {
			expos = append(expos, expo)
		}
	}

	return expos, nil
}

func AllCategories() ([][]string, error) {
	data, err := currentExpos()
	if err != nil {
		return nil, err
	}

	set := map[string]bool{}
	categories := [][]string{}
	for i, expo := range data {
		if i == 0 {
			continue
		}
		if _, ok := set[expo[4]]; !ok {
			set[expo[4]] = true
			c := []string{expo[3], expo[4]}
			categories = append(categories, c)
		}
	}

	return categories, nil
}

func Expos(tag string) ([][]string, error) {
	data, err := currentExpos()
	if err != nil {
		return nil, err
	}

	expos := [][]string{}

	for _, expo := range data {
		if expo[4] == tag || tag == "upcoming" {
			expos = append(expos, expo)
		}
	}

	return expos, nil
}
