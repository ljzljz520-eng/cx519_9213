package visual

import (
	"soundspace/internal/model"
	"testing"
)

func TestExportJSON(t *testing.T) {
	p := New()
	p.Apply(model.Record{ID: "x", Title: "X", Status: "draft"})
	b, e := p.ExportJSON()
	if e != nil || len(b) == 0 {
		t.Fatal(e)
	}
}
