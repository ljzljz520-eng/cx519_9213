package query

import (
	"sort"
	"soundspace/internal/model"
	"soundspace/internal/store"
)

type Engine struct{ Store *store.Store }

func New(s *store.Store) *Engine { return &Engine{Store: s} }
func (e *Engine) Find(f model.SearchFilter) (model.Page, error) {
	p, err := e.Store.Search(f)
	sort.SliceStable(p.Items, func(i, j int) bool { return p.Items[i].UpdatedAt.After(p.Items[j].UpdatedAt) })
	return p, err
}
func (e *Engine) Timeline(id string) ([]model.AuditEvent, error) { return nil, nil }
