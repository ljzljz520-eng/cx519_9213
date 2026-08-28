package store

import (
	"go.etcd.io/bbolt"
	"soundspace/internal/model"
	"strings"
)

func (s *Store) Search(f model.SearchFilter) (model.Page, error) {
	f = model.CheckFilter(f)
	rs, e := s.ListRecords()
	if e != nil {
		return model.Page{}, e
	}
	out := make([]model.Record, 0)
	for _, r := range rs {
		if f.Status != "" && r.Status != f.Status {
			continue
		}
		if f.Location != "" && !strings.EqualFold(r.Location, f.Location) {
			continue
		}
		if f.Text != "" && !strings.Contains(strings.ToLower(r.Title+" "+r.Notes), strings.ToLower(f.Text)) {
			continue
		}
		out = append(out, r)
	}
	total := len(out)
	if len(out) > f.Limit {
		out = out[:f.Limit]
	}
	return model.Page{Items: out, Total: total}, nil
}
func (s *Store) DeleteRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Delete([]byte(id)) })
}
func (s *Store) Count() (int, error) { rs, e := s.ListRecords(); return len(rs), e }
