package flow061

import "soundspace/internal/model"

func (s *Service) RegisterBatch(rows []model.Record, actor string) error {
	for _, r := range rows {
		if err := s.Register(r, actor); err != nil {
			return err
		}
	}
	return nil
}
func (s *Service) Snapshot(id string) (model.Record, error) { return s.Store.GetRecord(id) }
