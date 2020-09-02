package jobs

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	q := New()
	id := q.Enqueue(&EnqueueOpts{
		Type: NotTimeCritical,
	})
	job := q.Get(id)
	if id != offset {
		t.Fatalf("expected id to match offset: %d != %d", id, offset)
	}
	if id != job.ID {
		t.Fatalf("expected id to match job ID: %d != %d", id, job.ID)
	}
	if job.Status != Queued {
		t.Fatalf("expected status to be queued: %q != %q", job.Status, Queued)
	}
	if job.Type != NotTimeCritical {
		t.Fatalf("expected type to be NotTimeCritical: %q != %q", job.Type, NotTimeCritical)
	}
}

func TestMultipleEnqueues(t *testing.T) {
	q := New()
	first := q.Enqueue(&EnqueueOpts{
		Type: NotTimeCritical,
	})
	_ = q.Enqueue(&EnqueueOpts{
		Type: TimeCritical,
	})
	_ = q.Enqueue(&EnqueueOpts{
		Type: NotTimeCritical,
	})
	j := q.Dequeue()
	if j.Status != InProgress {
		t.Fatalf("expected dequeued job to be in progress: %q != %q", j.Status, InProgress)
	}
	q.Conclude(first)
	concluded := q.Get(first)
	if concluded.Status != Concluded {
		t.Fatalf("expected concluded job to be concluded: %q != %q", concluded.Status, Concluded)
	}
}
