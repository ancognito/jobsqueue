package api

import (
	"encoding/json"
	"net/http"

	"github.com/ancognito/jobsqueue/internal/jobs"
)

type DequeueRequest struct{}

type DequeueResponse struct {
	ID     jobs.ID     `json:"ID"`
	Type   jobs.Type   `json:"Type"`
	Status jobs.Status `json:"Status"`
}

func (h *V1Handler) DequeueHandlerV1(w http.ResponseWriter, r *http.Request) {
	p := &DequeueResponse{}

	job := h.jobs.Dequeue()
	if job == nil {
		// TODO - decide what API contract to support here
		w.WriteHeader(204)
		return
	}

	p.ID = job.ID
	p.Type = job.Type
	p.Status = job.Status

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
