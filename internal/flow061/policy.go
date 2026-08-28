package flow061

import "soundspace/internal/model"

func CanEdit(r model.Record) bool { return r.Status != "archived" }
func NextStages(status string) []string {
	switch status {
	case "draft":
		return []string{"review"}
	case "review":
		return []string{"approved", "draft"}
	case "approved":
		return []string{"archived", "review"}
	default:
		return []string{}
	}
}
