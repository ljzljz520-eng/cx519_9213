package encoding

import (
	"encoding/json"
	"fmt"
	"soundspace/internal/model"
	"strconv"
	"strings"
)

func EncodeRecord(r model.Record) []byte { b, _ := json.Marshal(r); return b }
func DecodeRecord(b []byte) (model.Record, error) {
	var r model.Record
	e := json.Unmarshal(b, &r)
	return r, e
}
func EncodeLine(r model.Record) string {
	return strings.Join([]string{r.ID, r.Title, r.Location, r.Frequency, r.Status, r.Notes}, "|")
}
func DecodeLine(line string) (model.Record, error) {
	p := strings.Split(line, "|")
	if len(p) != 6 {
		return model.Record{}, fmt.Errorf("fields")
	}
	return model.Record{ID: p[0], Title: p[1], Location: p[2], Frequency: p[3], Status: p[4], Notes: p[5]}, nil
}
func Int(v int) string               { return strconv.Itoa(v) }
func ParseInt(v string) (int, error) { return strconv.Atoi(strings.TrimSpace(v)) }
