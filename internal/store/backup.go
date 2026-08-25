package store

import (
	"encoding/json"
	"soundspace/internal/model"
)

func (s *Store) ExportRecords() ([]byte, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	return json.Marshal(rs)
}
func DecodeRecords(data []byte) ([]model.Record, error) {
	var rs []model.Record
	e := json.Unmarshal(data, &rs)
	return rs, e
}
func (s *Store) ImportRecords(rs []model.Record) error {
	for _, r := range rs {
		if e := s.SaveRecord(r); e != nil {
			return e
		}
	}
	return nil
}
