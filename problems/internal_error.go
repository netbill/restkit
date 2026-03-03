package problems

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/jsonapi"
)

func InternalError() error {
	return &jsonapi.ErrorObject{
		Title:  http.StatusText(http.StatusInternalServerError),
		Status: fmt.Sprintf("%d", http.StatusInternalServerError),
		Code:   "INTERNAL_SERVER_ERROR",
		Meta: &map[string]any{
			"timestamp": time.Now().UTC(),
		},
	}
}
