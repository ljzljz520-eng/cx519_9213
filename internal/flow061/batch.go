package flow061

import "soundspace/internal/model"

func (s *Service) ApproveBatch(ids []string, actor string) (int, []error) {
	n := 0
	errs := []error{}
	for _, id := range ids {
		if e := s.Approve(id, actor, "batch approval"); e != nil {
			errs = append(errs, e)
		} else {
			n++
		}
	}
	return n, errs
}
func (s *Service) ArchiveBatch(ids []string, actor string) (int, []error) {
	n := 0
	errs := []error{}
	for _, id := range ids {
		if e := s.Archive(id, actor); e != nil {
			errs = append(errs, e)
		} else {
			n++
		}
	}
	return n, errs
}
func BuildRecord(id, title, location, freq string) model.Record {
	return model.Record{ID: id, Title: title, Location: location, Frequency: freq, Status: "draft"}
}
