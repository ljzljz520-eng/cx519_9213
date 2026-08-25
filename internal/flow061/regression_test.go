package flow061

import (
	"soundspace/internal/model"
	"soundspace/internal/visual"
	"testing"
)

func Test519BusinessRegression(t *testing.T) {
	p := visual.New()
	a := p.Apply(model.Record{ID: "a", Title: "First", Status: "draft", Revision: 1})
	b := p.Apply(model.Record{ID: "b", Title: "Second", Status: "draft", Revision: 1})
	if a.State.Label != "First" {
		t.Fatal("first state")
	}
	if b.State.Label != "Second" {
		t.Fatalf("second state displayed %q", b.State.Label)
	}
}
