package api

import (
	"io"
	"net/http"

	"github.com/ancognito/jobsqueue/internal/jobs"
	"github.com/gorilla/mux"
)

type V1Handler struct {
	router *mux.Router
	jobs   *jobs.Queue
}

func (h *V1Handler) EmptyHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "")
}

func NewV1Handler(q *jobs.Queue) *V1Handler {
	r := mux.NewRouter()

	h := &V1Handler{
		router: r,
		jobs:   q,
	}

	r.Methods("GET").Path("/jobs/enqueue").HandlerFunc(h.EnqueueHandlerV1)
	r.Methods("POST").Path("/jobs/enqueue").HandlerFunc(h.EnqueueHandlerV1)
	r.Methods("GET").Path("/jobs/dequeue").HandlerFunc(h.DequeueHandlerV1)
	r.Methods("GET").Path("/jobs/{id:[0-9]+}/conclude").HandlerFunc(h.ConcludeHandlerV1)
	r.Methods("GET").Path("/jobs/{id:[0-9]+}").HandlerFunc(h.JobHandlerV1)

	return h
}

func (h *V1Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(res, req)
}
