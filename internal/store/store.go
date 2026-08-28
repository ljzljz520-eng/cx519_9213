package store

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"soundspace/internal/model"
	"sync"
	"time"
)

var buckets = [][]byte{[]byte("records"), []byte("audits"), []byte("workflows"), []byte("attachments")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func put(tx *bbolt.Tx, b []byte, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put([]byte(key), data)
}
func get(tx *bbolt.Tx, b []byte, key string, v any) error {
	data := tx.Bucket(b).Get([]byte(key))
	if data == nil {
		return fmt.Errorf("not found")
	}
	return json.Unmarshal(data, v)
}
func (s *Store) SaveRecord(r model.Record) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[0], r.ID, r) })
}
func (s *Store) GetRecord(id string) (model.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var r model.Record
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets[0], id, &r) })
	return r, e
}
func (s *Store) SaveAudit(a model.AuditEvent) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[1], a.ID, a) })
}
func (s *Store) SaveWorkflow(w model.Workflow) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[2], w.ID, w) })
}
func (s *Store) SaveAttachment(a model.Attachment) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[3], a.ID, a) })
}
func (s *Store) ListRecords() ([]model.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[0]).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
