package metrics

import (
	"sort"
	"soundspace/internal/model"
)

type Bucket struct {
	Name  string
	Count int
}

func SortedStatuses(rs []model.Record) []Bucket {
	m := StatusCounts(rs)
	out := make([]Bucket, 0, len(m))
	for k, v := range m {
		out = append(out, Bucket{k, v})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Count == out[j].Count {
			return out[i].Name < out[j].Name
		}
		return out[i].Count > out[j].Count
	})
	return out
}
func LocationCounts(rs []model.Record) map[string]int {
	out := map[string]int{}
	for _, r := range rs {
		out[r.Location]++
	}
	return out
}
func RevisionTotal(rs []model.Record) int {
	n := 0
	for _, r := range rs {
		n += r.Revision
	}
	return n
}
