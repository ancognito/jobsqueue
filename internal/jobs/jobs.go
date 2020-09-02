package jobs

import (
	"container/list"
)

const (
	offset = 100
)

// TODO - *MUST* add errors to this, intentionally skipped for now

// JobQueuer defines the contract expected by consumers of this Queue
type JobQueuer interface {
	Enqueue(*EnqueueOpts) ID
	Dequeue() *Job
	Conclude(ID)
	Get(ID) *Job
}

// Queue implements a FIFO queue with in-memory persistence
type Queue struct {
	waiting *list.List
	jobs    map[ID]*Job
	next    ID
}

// New returns an empty, initialized Queue
func New() *Queue {
	return &Queue{
		waiting: list.New(),
		jobs:    make(map[ID]*Job),
		next:    offset,
	}
}

// EnqueueOpts allows additional data to be provided when enqueueing a Job
type EnqueueOpts struct {
	Type Type
}

// Enqueue pushes a Job onto the end of the Queue
// it is not concurrency-safe!
func (q *Queue) Enqueue(m *EnqueueOpts) ID {
	j := &Job{
		ID:     q.next,
		Type:   m.Type,
		Status: Queued,
	}
	q.next++
	q.jobs[j.ID] = j
	q.waiting.PushBack(j.ID)
	return j.ID
}

// Dequeue removes a Job from the Queue, or
// it returns nil if the Queue is empty
func (q *Queue) Dequeue() *Job {
	if q.waiting.Len() <= 0 {
		return nil
	}
	id := q.waiting.Remove(
		q.waiting.Front(),
	).(ID)
	q.jobs[id].Status = InProgress
	return q.jobs[id]
}

// Conclude flags this job as Concluded
func (q *Queue) Conclude(id ID) {
	_, ok := q.jobs[id]
	if !ok {
		// TODO - appropriate error
		return
	}
	q.jobs[id].Status = Concluded
}

// Get returns the job by ID, or returns nil
func (q *Queue) Get(id ID) *Job {
	return q.jobs[id]
}
