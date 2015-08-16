package back

import (
	"log"
)

const (
	URL = "https://docs.google.com/spreadsheets/d/1ud6IZFjoT0Hyvh6TjbGq7czDezs9M3ODqEKysQ04h8E/pub?gid=0&single=true&output=csv"
)

func SiteCopy() ([][]string, error) {
	if data, err := readCSVFromUrl(URL); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return data, nil
	}

	return nil, ErrKeyNotFound
}
