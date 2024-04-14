package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

func slogMap(key string, value map[string]any) slog.Attr {
	out := make([]slog.Attr, 0)
	for k, v := range value {
		out = append(out, slog.Any(k, v))
	}
	return slog.Attr{Key: key, Value: slog.GroupValue(out...)}
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	_ = logger
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		payload := make(map[string]any)
		if err != nil {
			fmt.Printf("Error reading request body: %v\n", err)
		} else {
			if len(body) > 0 {
				err := json.Unmarshal(body, &payload)
				if err != nil {
					fmt.Printf("Error reading request body: %v\n", err)
				}
			}
			r.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		logger.Debug("REQUEST",
			slog.String("method", r.Method),
			slog.String("path", r.URL.EscapedPath()),
			slogMap("payload", payload),
		)
		next.ServeHTTP(w, r)
	})
}
