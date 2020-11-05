package main

import (
	"net/http"
	"runtime"
	"go.uber.org/zap"
	"encoding/json"
)

func (s *server) healthCheck(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key).([]byte)
	s.logger.Info("new health request",
		zap.String("request_ip", r.RemoteAddr),
		zap.Binary("request_id", id),
	)

	w.WriteHeader(http.StatusOK)
}

func (s *server) fullHealthCheck(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key).([]byte)
	s.logger.Info("new full health request",
		zap.String("request_ip", r.RemoteAddr),
		zap.Binary("request_id", id),
	)

	ms := &runtime.MemStats{}
	runtime.ReadMemStats(ms)

	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(ms)
	if err != nil {
		s.logger.Error("failed to encode JSON",
			zap.String("request_ip", r.RemoteAddr),
			zap.Error(err),
		)
		http.Error(w, "failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
