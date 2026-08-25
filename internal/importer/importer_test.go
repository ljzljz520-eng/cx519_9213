package importer

import "testing"

func TestParseReport(t *testing.T) {
	rows, r := Parse([]string{"a,Alpha,N,100", "bad"})
	if len(rows) != 1 || r.Accepted != 1 || r.Rejected != 1 {
		t.Fatalf("%#v %#v", rows, r)
	}
}
