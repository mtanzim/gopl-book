package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type XKCDSummary struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func GetComic(id string, c *Cache) (*XKCDSummary, string, error) {
	if result, ok := c.GetFromCache(id); ok {
		return result, "cache", nil
	}
	remoteResult, err := getComicFromRemote(id)
	if err != nil {
		return nil, "error", err
	}
	c.SetCache(id, remoteResult)
	return remoteResult, "remote", err
}

func makeURL(id string) string {
	return fmt.Sprintf("https://xkcd.com/%s/info.0.json", id)
}

func getComicFromRemote(id string) (*XKCDSummary, error) {
	url := makeURL(id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	if err != nil {
		return nil, err
	}
	var result XKCDSummary
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
