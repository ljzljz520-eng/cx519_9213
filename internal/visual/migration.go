package visual

import "soundspace/internal/model"

func (p *Processor) Migrate(records []model.Record) []Frame {
	out := make([]Frame, 0, len(records))
	for _, r := range records {
		out = append(out, p.Apply(r))
	}
	return out
}
func (p *Processor) Reset() { p.current = State{}; p.frames = nil }
func (p *Processor) Find(id string) (Frame, bool) {
	for _, f := range p.frames {
		if f.State.RecordID == id {
			return f, true
		}
	}
	return Frame{}, false
}
