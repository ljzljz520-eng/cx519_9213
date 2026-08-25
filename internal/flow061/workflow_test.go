package flow061

import (
	"path/filepath"
	"soundspace/internal/model"
	"soundspace/internal/store"
	"testing"
)

func TestWorkflowCreateReviewArchive(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "a.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	f := New(s)
	if e = f.Register(model.Record{ID: "r1", Title: "North", Location: "A", Status: "draft"}, "u"); e != nil {
		t.Fatal(e)
	}
	if e = f.Review("r1", "u"); e != nil {
		t.Fatal(e)
	}
	if e = f.Approve("r1", "u", "ok"); e != nil {
		t.Fatal(e)
	}
	if e = f.Archive("r1", "u"); e != nil {
		t.Fatal(e)
	}
	r, _ := s.GetRecord("r1")
	if r.Status != "archived" {
		t.Fatalf("status %s", r.Status)
	}
}
