package security

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type Principal struct {
	Name  string
	Roles []string
}

func NewPrincipal(name string, roles ...string) Principal {
	return Principal{Name: name, Roles: append([]string{}, roles...)}
}
func (p Principal) Has(role string) bool {
	for _, r := range p.Roles {
		if r == role {
			return true
		}
	}
	return false
}
func (p Principal) Can(action string) bool {
	switch action {
	case "register":
		return p.Has("analyst") || p.Has("admin")
	case "review":
		return p.Has("reviewer") || p.Has("admin")
	case "archive":
		return p.Has("curator") || p.Has("admin")
	default:
		return false
	}
}
func Digest(data string) string    { h := sha256.Sum256([]byte(data)); return hex.EncodeToString(h[:]) }
func CleanName(name string) string { return strings.TrimSpace(strings.ToLower(name)) }
