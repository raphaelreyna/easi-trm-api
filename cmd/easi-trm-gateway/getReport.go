package main

import (
	"net/http"
	trm "github.com/raphaelreyna/easi-trm-api"
	"encoding/json"
)

func (s *server) handleGetReport() http.HandlerFunc {
	// Define container structs for sending and receiving data with connected HTTP client
	type req struct {
		DevKey string `json:"DevKey"`
		trm.ReportRequest
	}
	type resp struct {
		EasiReportMessage string `json:"EasiReportMessage"`
		EasiReportFieldData []*trm.FieldData `json:"EasiReportFieldData"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode the request body into the DevKey and ReportRequest container req struct
		rr := &req{}
		if err := json.NewDecoder(r.Body).Decode(rr); err != nil {
			s.err(w, r, "failed to decode JSON", err)
			return
		}

		// Request the report from the EASI servers on behalf of the client
		data, msg, err := s.client.GetReportContext(r.Context(), &rr.ReportRequest, rr.DevKey, r.Header)
		if err != nil {
			s.err(w, r, "failed to get report data from EASI", err)
			return

		}

		// Write out response from servers to the client as JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(&resp{msg, data}); err != nil {
			s.err(w, r, "failed to encode JSON", err)
			return
		}
	}
}
