package problems

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/jsonapi"
)

func Gone(details string) error {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusGone),
		Status: fmt.Sprintf("%d", http.StatusGone),
		Code:   "GONE",
		Detail: details,
		Meta: &map[string]any{
			"timestamp": time.Now().UTC(),
		},
	}
}
