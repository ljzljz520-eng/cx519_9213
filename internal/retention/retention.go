package retention

import (
	"sort"
	"soundspace/internal/model"
	"time"
)

type Policy struct {
	MaxAge       time.Duration
	KeepArchived bool
}

func Expired(r model.Record, now time.Time, p Policy) bool {
	if p.KeepArchived && r.Status == "archived" {
		return false
	}
	return now.Sub(r.UpdatedAt) > p.MaxAge
}
func Candidates(rs []model.Record, now time.Time, p Policy) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if Expired(r, now, p) {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.Before(out[j].UpdatedAt) })
	return out
}
func Partition(rs []model.Record, now time.Time, p Policy) ([]model.Record, []model.Record) {
	keep := []model.Record{}
	drop := []model.Record{}
	for _, r := range rs {
		if Expired(r, now, p) {
			drop = append(drop, r)
		} else {
			keep = append(keep, r)
		}
	}
	return keep, drop
}
func Age(r model.Record, now time.Time) time.Duration { return now.Sub(r.UpdatedAt) }
