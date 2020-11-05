package main

import (
	"net/http"
	"go.uber.org/zap"
	"net"
	"strings"
	"context"
	"math/rand"
)

func contextWithID(ctx context.Context, c net.Conn) context.Context {
	// Generate an id for this new request
	id := make([]byte, 8)
	_, err := rand.Read(id)
	if err != nil {
		panic(err)
		return context.Background()
	}

	return context.WithValue(ctx, key, id)
}

func (s *server) internalRoute(next http.HandlerFunc) http.HandlerFunc {
	if s.internalNet == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			next(w, r)
		}

	}
	return func(w http.ResponseWriter, r *http.Request) {
		ip := strings.Split(r.RemoteAddr, ":")[0]
		if !s.internalNet.Contains(net.ParseIP(ip)) {
			id := r.Context().Value(key).([]byte)
			s.logger.Info("internal route requested outside of internal network",
				zap.String("request_ip", r.RemoteAddr),
				zap.Binary("request_id", id),
				zap.String("internal_network", s.internalNet.String()),
			)
			http.NotFound(w, r)
			return
		}

		next(w, r)
	}
}

func (s *server) jsonRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(key).([]byte)
		headers := r.Header
		// Make "sure" the client is sending JSON
		if ct := headers.Get("Content-Type"); !strings.Contains(ct, "application/json") && ct != "" {
			s.logger.Info("received GetReport request with invalid Content-Type header",
				zap.String("request_ip", r.RemoteAddr),
				zap.Binary("request_id", id),
				zap.String("content_type", ct),
			)
			http.Error(w, "invalid Content-Type header: expected application/json", http.StatusUnsupportedMediaType)
			return
		}
		// Make sure the client is accepting JSON
		acceptsJSON := len(headers.Values("Accept")) == 0
		for _, h := range headers.Values("Accept") {
			if strings.Contains(h, "application/json") || strings.Contains(h, "*/*"){
				acceptsJSON = true
			}
		}
		if !acceptsJSON {
			s.logger.Info("received GetReport request with invalid Accept header",
				zap.String("request_ip", r.RemoteAddr),
				zap.Binary("request_id", id),
			)
			http.Error(w, "invalid Accept header: expected application/json", http.StatusNotAcceptable)
			return
		}

		next(w, r)
	}
}

type lrw struct {
	http.ResponseWriter
	statusCode int
}

func (w *lrw) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.statusCode = code
}

func (s *server) logRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(key).([]byte)

		s.logger.Info("begin request",
			zap.String("request_ip", r.RemoteAddr),
			zap.Binary("request_id", id),
			zap.String("request_url", r.RequestURI),
		)

		lw := &lrw{}
		lw.ResponseWriter = w

		next(lw, r)

		s.logger.Info("end request",
			zap.String("request_ip", r.RemoteAddr),
			zap.Binary("request_id", id),
			zap.String("request_url", r.URL.String()),
			zap.Int("response_status", lw.statusCode),
		)
	}
}
