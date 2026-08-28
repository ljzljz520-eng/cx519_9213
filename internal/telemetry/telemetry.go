package telemetry

import (
	"sort"
	"sync"
	"time"
)

type Event struct {
	Name  string
	At    time.Time
	Value int
}
type Sink struct {
	mu     sync.RWMutex
	events []Event
}

func New() *Sink { return &Sink{events: []Event{}} }
func (s *Sink) Record(name string, at time.Time, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, Event{name, at, value})
}
func (s *Sink) Snapshot() []Event {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := append([]Event{}, s.events...)
	sort.Slice(out, func(i, j int) bool { return out[i].At.Before(out[j].At) })
	return out
}
func (s *Sink) Count(name string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	n := 0
	for _, e := range s.events {
		if e.Name == name {
			n++
		}
	}
	return n
}
func (s *Sink) Clear() { s.mu.Lock(); defer s.mu.Unlock(); s.events = nil }
