package model

import "time"

type Record struct {
	ID, Title, Location, Frequency, Status, Notes string
	Revision                                      int
	CreatedAt, UpdatedAt                          time.Time
}
type AuditEvent struct {
	ID, RecordID, Actor, Action, Detail string
	At                                  time.Time
}
type Workflow struct {
	ID, RecordID, Stage, Owner string
	Version                    int
	UpdatedAt                  time.Time
}
type Attachment struct {
	ID, RecordID, Name, MediaType string
	Size                          int
	Digest                        string
}
type SearchFilter struct {
	Text, Status, Location string
	Limit                  int
}
type Page struct {
	Items []Record
	Total int
	Next  string
}
type Transition struct{ RecordID, From, To, Actor, Reason string }
type ImportRow struct{ ID, Title, Location, Frequency, Notes string }
type ImportReport struct {
	Accepted, Rejected int
	Errors             []string
}
