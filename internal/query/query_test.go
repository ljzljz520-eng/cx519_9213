package query

import (
	"path/filepath"
	"soundspace/internal/model"
	"soundspace/internal/store"
	"testing"
)

func TestSearchFilters(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "q.db"))
	defer s.Close()
	s.SaveRecord(model.Record{ID: "1", Title: "Alpha", Location: "North", Status: "draft"})
	s.SaveRecord(model.Record{ID: "2", Title: "Beta", Location: "South", Status: "review"})
	p, e := New(s).Find(model.SearchFilter{Location: "North"})
	if e != nil || len(p.Items) != 1 {
		t.Fatalf("%v %#v", e, p)
	}
}
