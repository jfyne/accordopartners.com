package back

import (
	"log"
)

const (
	SC_URL = "https://docs.google.com/spreadsheets/d/1ud6IZFjoT0Hyvh6TjbGq7czDezs9M3ODqEKysQ04h8E/pub?gid=0&single=true&output=csv"
)

func SiteCopy() ([][]string, error) {
	data, err := readCSVFromUrl(SC_URL)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}
