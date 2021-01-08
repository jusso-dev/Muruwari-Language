package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SearchResponse is struct for MeiliSearch result
type SearchResponse struct {
	Hits []struct {
		ID          string `json:"ID"`
		WordPhrase  string `json:"Word/Phrase"`
		Translation string `json:"Translation"`
	} `json:"hits"`
	Offset           int    `json:"offset"`
	Limit            int    `json:"limit"`
	NbHits           int    `json:"nbHits"`
	ExhaustiveNbHits bool   `json:"exhaustiveNbHits"`
	ProcessingTimeMs int    `json:"processingTimeMs"`
	Query            string `json:"query"`
}

// SearchPhrase accepts phrase or word to search by and returns indexed result from
// MeiliSearch
func SearchPhrase(phrase string) (SearchResponse, error) {

	url := fmt.Sprintf("http://192.168.1.50:7700/indexes/translations/search?q=%s&limit=1", phrase)

	searchRes := &SearchResponse{}
	err := getJSON(url, searchRes)
	if (err) != nil {
		return *searchRes, err
	}

	return *searchRes, nil
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {

	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
