package render

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/jsonapi"
)

// SSEHeaders sets the standard response headers for a Server-Sent Events
// stream. Call it once, before the first WriteSSE/WriteRawSSE/WriteErrorSSE.
func SSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}

// WriteSSE marshals v to JSON and writes it as one Server-Sent Events frame.
//
// Unlike Response/ResponseError, this returns an error instead of panicking:
// on a long-lived stream a write failure (the client disconnected) is a
// routine occurrence, not a bug, and callers are expected to log it
// themselves rather than have the request goroutine's recover() kick in.
func WriteSSE(w io.Writer, event string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal sse payload: %w", err)
	}
	return WriteRawSSE(w, event, data)
}

// WriteRawSSE writes a single SSE frame, forwarding an already-marshaled
// JSON payload as-is — e.g. a payload published by another request/replica.
func WriteRawSSE(w io.Writer, event string, data []byte) error {
	_, err := fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, data)
	return err
}

// WriteErrorSSE writes err as a single SSE frame, using the same JSON:API
// {"errors": [...]} envelope ResponseError writes for regular error
// responses. err must wrap a *jsonapi.ErrorObject (e.g. one built by the
// problems package).
func WriteErrorSSE(w io.Writer, event string, err error) error {
	var jo *jsonapi.ErrorObject
	if !errors.As(err, &jo) || jo == nil {
		return fmt.Errorf("write error sse: %w is not a jsonapi error", err)
	}

	data, marshalErr := json.Marshal(jsonapi.ErrorsPayload{Errors: []*jsonapi.ErrorObject{jo}})
	if marshalErr != nil {
		return fmt.Errorf("marshal error sse payload: %w", marshalErr)
	}

	return WriteRawSSE(w, event, data)
}
