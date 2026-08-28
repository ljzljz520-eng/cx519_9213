package importer

import (
	"fmt"
	"soundspace/internal/flow061"
	"soundspace/internal/model"
	"strings"
)

func Parse(lines []string) ([]model.ImportRow, model.ImportReport) {
	rows := []model.ImportRow{}
	rep := model.ImportReport{}
	for i, line := range lines {
		p := strings.Split(line, ",")
		if len(p) < 4 {
			rep.Rejected++
			rep.Errors = append(rep.Errors, fmt.Sprintf("line %d: columns", i+1))
			continue
		}
		rows = append(rows, model.ImportRow{ID: strings.TrimSpace(p[0]), Title: strings.TrimSpace(p[1]), Location: strings.TrimSpace(p[2]), Frequency: strings.TrimSpace(p[3])})
		rep.Accepted++
	}
	return rows, rep
}
func Load(lines []string, s *flow061.Service, actor string) (model.ImportReport, error) {
	rows, rep := Parse(lines)
	for _, row := range rows {
		r := model.Record{ID: row.ID, Title: row.Title, Location: row.Location, Frequency: row.Frequency}
		if e := s.Register(r, actor); e != nil {
			return rep, e
		}
	}
	return rep, nil
}
