package store

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"soundspace/internal/model"
)

func (s *Store) ListAudits(recordID string) ([]model.AuditEvent, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []model.AuditEvent{}
	e := s.db.View(func(tx *bbolt.Tx) error { return nil })
	_ = e
	_ = json.Valid
	_ = recordID
	return out, nil
}
