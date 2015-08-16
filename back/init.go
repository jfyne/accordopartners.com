package back

import (
	"encoding/csv"
	"errors"
	"github.com/pmylund/go-cache"
	"net/http"
	"time"
)

var (
	c              = cache.New(1*time.Hour, 10*time.Minute)
	ErrKeyNotFound = errors.New("Data key not found in output")
)

func readCSVFromUrl(url string) ([][]string, error) {
	if data, found := c.Get(url); found {
		return data.([][]string), nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	go c.Set(url, data, cache.DefaultExpiration)

	return data, nil
}
