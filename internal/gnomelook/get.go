package gnomelook

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pablodz/gnome-theme-manager/utils"
)

const URL_SCRAPPED_JSON = "https://raw.githubusercontent.com/pablodz/gnomelookscraper/main/data/result_only_with_sources.json"

func GetDataFromScraper() error {

	rBody, rCode, err := utils.HTTPRequest(&utils.RequestModel{
		URL:    URL_SCRAPPED_JSON,
		Method: http.MethodGet,
	})
	if err != nil {
		return err
	}

	if rCode != http.StatusOK {
		return fmt.Errorf("error getting data from scraper: %d", rCode)
	}

	// log.Println(rBody)

	err = json.Unmarshal([]byte(rBody), &DATA_THEMES)
	if err != nil {
		return err
	}

	return nil

}

func GetItems() ([]string, []string, []string) {

	longRepos := []string{}
	shortRepos := []string{}
	zipUrls := []string{}
	for _, theme := range DATA_THEMES {
		fullRepo := theme.Sources.Repository
		longRepos = append(longRepos, fullRepo)
		zipUrls = append(zipUrls, theme.Sources.ZipURL)
		sR := fullRepo[strings.LastIndex(fullRepo, "/")+1:]
		shortRepos = append(shortRepos, sR)
	}

	return longRepos, shortRepos, zipUrls
}
