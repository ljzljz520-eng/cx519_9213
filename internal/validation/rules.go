package validation

import (
	"fmt"
	"soundspace/internal/model"
	"strings"
)

type Issue struct{ Field, Message string }

func RecordIssues(r model.Record) []Issue {
	out := []Issue{}
	if strings.TrimSpace(r.ID) == "" {
		out = append(out, Issue{"id", "required"})
	}
	if len(r.Title) > 120 {
		out = append(out, Issue{"title", "too long"})
	}
	if r.Frequency == "" {
		out = append(out, Issue{"frequency", "required"})
	}
	if r.Revision < 0 {
		out = append(out, Issue{"revision", "negative"})
	}
	return out
}
func Valid(r model.Record) bool { return len(RecordIssues(r)) == 0 }
func Explain(r model.Record) error {
	is := RecordIssues(r)
	if len(is) == 0 {
		return nil
	}
	return fmt.Errorf("%s: %s", is[0].Field, is[0].Message)
}
func Unique(rs []model.Record) []Issue {
	seen := map[string]bool{}
	out := []Issue{}
	for _, r := range rs {
		if seen[r.ID] {
			out = append(out, Issue{r.ID, "duplicate"})
		}
		seen[r.ID] = true
	}
	return out
}
