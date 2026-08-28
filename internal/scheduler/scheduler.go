package scheduler

import (
	"fmt"
	"sort"
	"soundspace/internal/model"
	"time"
)

type Job struct {
	ID, RecordID, Kind string
	RunAt              time.Time
	Done               bool
}
type Queue struct{ jobs map[string]Job }

func New() *Queue { return &Queue{jobs: map[string]Job{}} }
func (q *Queue) Add(j Job) error {
	if j.ID == "" || j.RecordID == "" {
		return fmt.Errorf("identity")
	}
	q.jobs[j.ID] = j
	return nil
}
func (q *Queue) Due(now time.Time) []Job {
	out := []Job{}
	for _, j := range q.jobs {
		if !j.Done && !j.RunAt.After(now) {
			out = append(out, j)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].RunAt.Before(out[j].RunAt) })
	return out
}
func (q *Queue) Complete(id string) bool {
	j, ok := q.jobs[id]
	if !ok {
		return false
	}
	j.Done = true
	q.jobs[id] = j
	return true
}
func (q *Queue) Pending() int {
	n := 0
	for _, j := range q.jobs {
		if !j.Done {
			n++
		}
	}
	return n
}

var _ model.Record
