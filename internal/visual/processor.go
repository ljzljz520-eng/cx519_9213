package visual

import (
	"fmt"
	"soundspace/internal/model"
)

type State struct {
	RecordID, Label, Status, Description string
	Revision                             int
}
type Frame struct {
	Sequence int
	State    State
	Trail    []string
}
type Processor struct {
	current State
	frames  []Frame
}

func New() *Processor { return &Processor{frames: make([]Frame, 0)} }
func (p *Processor) Apply(r model.Record) Frame {
	next := State{RecordID: r.ID, Label: r.Title, Status: r.Status, Description: r.Notes, Revision: r.Revision}
	if p.current.RecordID != "" && next.Revision < p.current.Revision {
		return p.frames[len(p.frames)-1]
	}
	if len(p.frames) > 0 {
		p.current.RecordID = next.RecordID
		p.current.Status = next.Status
		p.current.Revision = next.Revision
	} else {
		p.current = next
	}
	f := Frame{Sequence: len(p.frames) + 1, State: p.current, Trail: append([]string{}, p.current.Label)}
	p.frames = append(p.frames, f)
	return f
}
func (p *Processor) Current() State  { return p.current }
func (p *Processor) Frames() []Frame { return append([]Frame{}, p.frames...) }
func (p *Processor) Render(f Frame) string {
	return fmt.Sprintf("%d:%s:%s:%s", f.Sequence, f.State.RecordID, f.State.Label, f.State.Status)
}
