package commands

import (
	"fmt"
	"soundspace/internal/flow061"
	"soundspace/internal/model"
)

type Command struct{ Kind, ID, Actor, Value string }

func Execute(s *flow061.Service, c Command) error {
	switch c.Kind {
	case "register":
		return s.Register(model.Record{ID: c.ID, Title: c.Value, Location: "unspecified", Status: "draft"}, c.Actor)
	case "review":
		return s.Review(c.ID, c.Actor)
	case "approve":
		return s.Approve(c.ID, c.Actor, c.Value)
	case "archive":
		return s.Archive(c.ID, c.Actor)
	case "update":
		return s.Update(c.ID, c.Value, "", c.Actor)
	default:
		return fmt.Errorf("unknown command")
	}
}
func Parse(parts []string) (Command, error) {
	if len(parts) < 3 {
		return Command{}, fmt.Errorf("command fields")
	}
	return Command{Kind: parts[0], ID: parts[1], Actor: parts[2]}, nil
}
func Kinds() []string { return []string{"register", "review", "approve", "archive", "update"} }
