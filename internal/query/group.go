package query

import (
	"soundspace/internal/model"
	"strings"
)

func GroupByLocation(rs []model.Record) map[string][]model.Record {
	out := map[string][]model.Record{}
	for _, r := range rs {
		k := strings.TrimSpace(r.Location)
		out[k] = append(out[k], r)
	}
	return out
}
func Titles(rs []model.Record) []string {
	out := make([]string, 0, len(rs))
	for _, r := range rs {
		out = append(out, r.Title)
	}
	return out
}
func Match(r model.Record, text string) bool {
	return text == "" || strings.Contains(strings.ToLower(r.Title+" "+r.Notes), strings.ToLower(text))
}
