package xkcd

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
