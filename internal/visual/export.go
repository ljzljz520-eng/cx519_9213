package visual

import (
	"encoding/json"
	"sort"
	"soundspace/internal/model"
)

type Export struct {
	RecordID string `json:"record_id"`
	Label    string `json:"label"`
	Status   string `json:"status"`
	Sequence int    `json:"sequence"`
}

func (p *Processor) ExportJSON() ([]byte, error) { items := p.Exports(); return json.Marshal(items) }
func (p *Processor) Exports() []Export {
	fs := p.Frames()
	out := make([]Export, 0, len(fs))
	for _, f := range fs {
		out = append(out, Export{RecordID: f.State.RecordID, Label: f.State.Label, Status: f.State.Status, Sequence: f.Sequence})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Sequence < out[j].Sequence })
	return out
}
func ToState(r model.Record) State {
	return State{RecordID: r.ID, Label: r.Title, Status: r.Status, Description: r.Notes, Revision: r.Revision}
}
