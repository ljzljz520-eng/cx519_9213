package review

import (
	"fmt"
	"soundspace/internal/model"
	"soundspace/internal/store"
)

type Queue struct{ Store *store.Store }

func New(s *store.Store) *Queue { return &Queue{Store: s} }
func (q *Queue) Check(r model.Record) error {
	if r.Status != "review" {
		return fmt.Errorf("not reviewable")
	}
	return nil
}
func (q *Queue) Decision(id, actor, reason string, approve bool) error {
	r, e := q.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if e = q.Check(r); e != nil {
		return e
	}
	if approve {
		r.Status = "approved"
	} else {
		r.Status = "draft"
	}
	r.Revision++
	return q.Store.SaveRecord(r)
}
