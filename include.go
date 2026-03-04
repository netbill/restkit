package restkit

import (
	"net/http"
	"slices"
	"strings"
)

func ParseIncludes(r *http.Request) []string {
	includes := make([]string, 0)
	for _, v := range r.URL.Query()["include"] {
		for _, part := range strings.Split(v, ",") {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			if !slices.Contains(includes, part) {
				includes = append(includes, part)
			}
		}
	}
	return includes
}
