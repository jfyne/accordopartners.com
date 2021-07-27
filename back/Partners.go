package back

import (
	"log"
)

const (
	P_URL = "https://docs.google.com/spreadsheets/d/1ud6IZFjoT0Hyvh6TjbGq7czDezs9M3ODqEKysQ04h8E/pub?gid=3&single=true&output=csv"
)

func Partners() ([][]string, error) {
	data, err := readCSVFromUrl(P_URL)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return data, nil
}
