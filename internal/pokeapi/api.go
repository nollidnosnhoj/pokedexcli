package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Client) GetLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	url := baseUrl + "/location-area?offset=0&limit=20"
	if pageUrl != nil {
		url = *pageUrl
	}

	cacheKey, err := getKeyFromUrl(url)
	if err != nil {
		return LocationAreasResponse{}, errors.New("failed to parse page url")
	}
	cacheEntry, cacheFound := c.cache.Get(cacheKey)

	if cacheFound {
		var res LocationAreasResponse
		err := json.Unmarshal(cacheEntry, &res)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return res, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad response: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Set(cacheKey, data)

	var res LocationAreasResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return res, nil
}

func getKeyFromUrl(urlVal string) (string, error) {
	parsedUrl, err := url.Parse(urlVal)
	if err != nil {
		return "", err
	}
	query := parsedUrl.RawQuery;
	if query == "" {
		return "location-area", nil
	}
	return "location-area?" + parsedUrl.RawQuery, nil
}