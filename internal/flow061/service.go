package flow061

import (
	"fmt"
	"soundspace/internal/model"
	"soundspace/internal/store"
	"time"
)

type Service struct {
	Store *store.Store
	clock func() time.Time
}

func New(s *store.Store) *Service {
	return &Service{Store: s, clock: func() time.Time { return time.Unix(1700000000, 0) }}
}
func (s *Service) Register(r model.Record, actor string) error {
	r.Status = model.NormalizeStatus(r.Status)
	if err := r.Validate(); err != nil {
		return err
	}
	now := s.clock()
	r.CreatedAt = now
	r.UpdatedAt = now
	r.Revision = 1
	if err := s.Store.SaveRecord(r); err != nil {
		return err
	}
	return s.audit(r.ID, actor, "register", "record created")
}
func (s *Service) Review(id, actor string) error {
	return s.transition(id, "review", actor, "submitted for review")
}
func (s *Service) Approve(id, actor, reason string) error {
	return s.transition(id, "approved", actor, reason)
}
func (s *Service) Archive(id, actor string) error {
	return s.transition(id, "archived", actor, "retired")
}
func (s *Service) transition(id, to, actor, reason string) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	t := model.Transition{RecordID: id, From: r.Status, To: to, Actor: actor, Reason: reason}
	if e = model.ValidateTransition(t); e != nil {
		return e
	}
	r.Status = to
	r.Revision++
	r.UpdatedAt = s.clock()
	if e = s.Store.SaveRecord(r); e != nil {
		return e
	}
	return s.audit(id, actor, "transition", to)
}
func (s *Service) audit(id, actor, action, detail string) error {
	a := model.AuditEvent{ID: fmt.Sprintf("%s-%s-%d", id, action, s.clock().UnixNano()), RecordID: id, Actor: actor, Action: action, Detail: detail, At: s.clock()}
	return s.Store.SaveAudit(a)
}
func (s *Service) Update(id, title, notes, actor string) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status == "archived" {
		return fmt.Errorf("archived record")
	}
	if title != "" {
		r.Title = title
	}
	r.Notes = notes
	r.Revision++
	r.UpdatedAt = s.clock()
	if e = s.Store.SaveRecord(r); e != nil {
		return e
	}
	return s.audit(id, actor, "update", "metadata changed")
}
func (s *Service) Publish(id, actor string) error {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return e
	}
	if r.Status != "approved" {
		return fmt.Errorf("approval required")
	}
	return s.audit(id, actor, "publish", "published")
}
