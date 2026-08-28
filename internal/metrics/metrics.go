package metrics

import "soundspace/internal/model"

func StatusCounts(rs []model.Record) map[string]int {
	out := map[string]int{}
	for _, r := range rs {
		out[r.Status]++
	}
	return out
}
func Frequencies(rs []model.Record) map[string]int {
	out := map[string]int{}
	for _, r := range rs {
		out[r.Frequency]++
	}
	return out
}
func Completion(rs []model.Record) float64 {
	if len(rs) == 0 {
		return 0
	}
	n := 0
	for _, r := range rs {
		if r.Status == "archived" {
			n++
		}
	}
	return float64(n) / float64(len(rs))
}
