package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Status  int         `json:"status,omitempty"`
	Success bool        `json:"success,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func New(log *slog.Logger, w *http.ResponseWriter, r *http.Request, data interface{}, statusCode int) {
	l := log.With("method", "response")
	(*w).WriteHeader(statusCode)

	switch v := data.(type) {
	case error:
		response := &Response{
			Status: statusCode,
			Error:  v.Error(),
		}

		data, err := json.Marshal(response)
		if err != nil {
			l.Error("error", err)
		}

		(*w).Write(data)
	case bool:
		response := &Response{
			Success: data.(bool),
		}

		data, err := json.Marshal(response)
		if err != nil {
			l.Error("error", err)
		}

		(*w).Write(data)

	default:
		response := &Response{
			Status: statusCode,
			Data:   v,
		}

		data, err := json.Marshal(response)
		if err != nil {
			l.Error("error", err)
		}

		(*w).Write(data)
	}
}
