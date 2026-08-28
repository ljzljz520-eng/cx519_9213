package report

import (
	"fmt"
	"soundspace/internal/model"
	"strings"
)

func Summary(page model.Page) string {
	var b strings.Builder
	fmt.Fprintf(&b, "total=%d\n", page.Total)
	for _, r := range page.Items {
		fmt.Fprintf(&b, "%s|%s|%s\n", r.ID, r.Title, r.Status)
	}
	return b.String()
}
func CSV(page model.Page) string {
	var b strings.Builder
	b.WriteString("id,title,status\n")
	for _, r := range page.Items {
		fmt.Fprintf(&b, "%s,%s,%s\n", r.ID, r.Title, r.Status)
	}
	return b.String()
}
