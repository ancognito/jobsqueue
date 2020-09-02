package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ancognito/jobsqueue/internal/jobs"
)

func (h *V1Handler) ConcludeHandlerV1(w http.ResponseWriter, r *http.Request) {
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

	h.jobs.Conclude(jobs.ID(id))

	w.WriteHeader(204)
}
