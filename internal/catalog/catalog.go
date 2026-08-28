package catalog

import (
	"sort"
	"soundspace/internal/model"
	"strings"
)

type Catalog struct{ items map[string]model.Record }

func New() *Catalog { return &Catalog{items: map[string]model.Record{}} }
func (c *Catalog) Add(r model.Record) error {
	if e := r.Validate(); e != nil {
		return e
	}
	c.items[r.ID] = r
	return nil
}
func (c *Catalog) Remove(id string) bool {
	if _, ok := c.items[id]; !ok {
		return false
	}
	delete(c.items, id)
	return true
}
func (c *Catalog) Get(id string) (model.Record, bool) { r, ok := c.items[id]; return r, ok }
func (c *Catalog) All() []model.Record {
	out := make([]model.Record, 0, len(c.items))
	for _, r := range c.items {
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
func (c *Catalog) ByStatus(status string) []model.Record {
	out := []model.Record{}
	for _, r := range c.items {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out
}
func (c *Catalog) ByText(text string) []model.Record {
	out := []model.Record{}
	text = strings.ToLower(text)
	for _, r := range c.items {
		if strings.Contains(strings.ToLower(r.Title), text) || strings.Contains(strings.ToLower(r.Notes), text) {
			out = append(out, r)
		}
	}
	return out
}
func (c *Catalog) Replace(r model.Record) bool {
	if _, ok := c.items[r.ID]; !ok {
		return false
	}
	c.items[r.ID] = r
	return true
}
func (c *Catalog) Size() int { return len(c.items) }
func (c *Catalog) Clear()    { c.items = map[string]model.Record{} }
