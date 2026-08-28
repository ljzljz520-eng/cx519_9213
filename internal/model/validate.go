package model

import "fmt"

func (r Record) Validate() error {
	if r.ID == "" {
		return fmt.Errorf("id required")
	}
	if r.Title == "" {
		return fmt.Errorf("title required")
	}
	if r.Location == "" {
		return fmt.Errorf("location required")
	}
	if r.Status == "" {
		return fmt.Errorf("status required")
	}
	return nil
}
func ValidateTransition(t Transition) error {
	allowed := map[string]map[string]bool{"draft": {"review": true}, "review": {"approved": true, "draft": true}, "approved": {"archived": true, "review": true}, "archived": {}}
	if _, ok := allowed[t.From]; !ok {
		return fmt.Errorf("unknown source")
	}
	if !allowed[t.From][t.To] {
		return fmt.Errorf("invalid transition %s to %s", t.From, t.To)
	}
	if t.RecordID == "" || t.Actor == "" {
		return fmt.Errorf("transition identity required")
	}
	return nil
}
func NormalizeStatus(s string) string {
	if s == "" {
		return "draft"
	}
	switch s {
	case "draft", "review", "approved", "archived":
		return s
	default:
		return "draft"
	}
}
func CheckFilter(f SearchFilter) SearchFilter {
	if f.Limit <= 0 || f.Limit > 100 {
		f.Limit = 25
	}
	return f
}
