package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ancognito/jobsqueue/internal/jobs"
)

type JobResponse struct {
	ID     jobs.ID     `json:"ID"`
	Type   jobs.Type   `json:"Type"`
	Status jobs.Status `json:"Status"`
}

func (h *V1Handler) JobHandlerV1(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	val, ok := params["id"]
	if !ok {
		http.Error(w, "unable to parse id from path", http.StatusBadRequest)
	}
	id, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("error converting id to int: %v", err)
		http.Error(w, "unable to parse id from path", http.StatusBadRequest)
	}

	p := &JobResponse{}
	job := h.jobs.Get(jobs.ID(id))

	if job == nil {
		// TODO - decide what API contract to support here
		http.Error(w, fmt.Sprintf("job %d not found", id), http.StatusNotFound)
		return
	}

	p.ID = job.ID
	p.Type = job.Type
	p.Status = job.Status

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}
