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

type Cache struct {
	Values map[string]*XKCDSummary
}

func NewCache() *Cache {
	return &Cache{Values: make(map[string]*XKCDSummary)}
}

func (c *Cache) SetCache(id string, result *XKCDSummary) {
	c.Values[id] = result
}

func (c *Cache) GetFromCache(id string) (*XKCDSummary, bool) {
	val, ok := c.Values[id]
	return val, ok
}

func makeURL(id string) string {
	return fmt.Sprintf("https://xkcd.com/%s/info.0.json", id)

}

// func GetComic(id string)(*XKCDSummary, error) {
// 	if result, ok := Ge
// }

func GetComicFromRemote(id string) (*XKCDSummary, error) {
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
