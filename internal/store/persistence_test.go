package store

import (
	"path/filepath"
	"soundspace/internal/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "store.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveRecord(model.Record{ID: "persist", Title: "Saved", Location: "L", Status: "draft"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, e := s.GetRecord("persist")
	if e != nil || r.Title != "Saved" {
		t.Fatalf("%v %#v", e, r)
	}
}
