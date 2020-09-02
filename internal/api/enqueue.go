package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ancognito/jobsqueue/internal/jobs"
)

type EnqueueRequest struct {
	Type jobs.Type `json:"Type"`
}

type EnqueueResponse struct {
	ID jobs.ID `json:"ID"`
}

func (h *V1Handler) EnqueueHandlerV1(w http.ResponseWriter, r *http.Request) {
	var reqBody EnqueueRequest
	var p EnqueueResponse
	var err error

	switch r.Method {
	case "GET":
		// just a convenience for testing
		reqBody.Type = jobs.NotTimeCritical
	case "POST":
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			log.Printf("unable to parse request: %v", err)
			http.Error(w, "unable to parse request", http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		return
	}

	if err != nil {
		// normally we wouldn't expose errors:
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	p.ID = h.jobs.Enqueue(&jobs.EnqueueOpts{
		Type: reqBody.Type,
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
