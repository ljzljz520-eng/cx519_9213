package model

import "strings"

func (r Record) Key() string         { return strings.ToLower(strings.TrimSpace(r.ID)) }
func (r Record) IsTerminal() bool    { return r.Status == "archived" }
func (r Record) Copy() Record        { return r }
func (a AuditEvent) Summary() string { return a.Action + ":" + a.Detail }
func (w Workflow) Active() bool      { return w.Stage != "archived" }
func (a Attachment) Valid() bool     { return a.ID != "" && a.RecordID != "" && a.Size >= 0 }
