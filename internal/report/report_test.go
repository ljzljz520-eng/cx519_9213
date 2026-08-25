package report

import (
	"soundspace/internal/model"
	"strings"
	"testing"
)

func TestCSV(t *testing.T) {
	s := CSV(model.Page{Items: []model.Record{{ID: "1", Title: "A", Status: "draft"}}})
	if !strings.Contains(s, "id,title,status") || !strings.Contains(s, "1,A,draft") {
		t.Fatal(s)
	}
}
