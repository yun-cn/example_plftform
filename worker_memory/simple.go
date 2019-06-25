package worker

import (
	"sync/atomic"
	"time"

	buffalo "github.com/gobuffalo/buffalo/worker"
)

type (
	// Job alias
	Job = buffalo.Job
	// Handler alias
	Handler = buffalo.Handler
	// Args alias
	Args = buffalo.Args
)

// SimpleWorker is a wrapper around worker.Simple that keeps track of it's job count so it can be shutdown gracefully without loosing any jobs.
type SimpleWorker struct {
	*buffalo.Simple
	JobsInProgress *int64
}

// Worker is the main noonde worker
var Worker *SimpleWorker

func init() {
	Worker = NewSimpleWorker()
}

// NewSimpleWorker ..
func NewSimpleWorker() *SimpleWorker {
	count := int64(0)
	return &SimpleWorker{buffalo.NewSimple(), &count}
}

// Increment job count
func (w *SimpleWorker) Increment() {
	atomic.AddInt64(w.JobsInProgress, 1)
}

// Decrement job count
func (w *SimpleWorker) Decrement() {
	atomic.AddInt64(w.JobsInProgress, -1)
}

// Register wraps the handler to decrement job count then delegates to buffalo.Simple
func (w SimpleWorker) Register(key string, h Handler) error {
	wrappedHandler := func(args Args) error {
		err := h(args)
		w.Decrement()
		return err
	}

	return w.Simple.Register(key, wrappedHandler)
}

// Perform increments count and delegates to buffalo.Simple
func (w SimpleWorker) Perform(j Job) error {
	w.Increment()
	return w.Simple.Perform(j)
}

// PerformAt increments count and delegates to buffalo.Simple
func (w SimpleWorker) PerformAt(j Job, t time.Time) error {
	w.Increment()
	return w.Simple.PerformAt(j, t)
}

// PerformIn increments count and delegates to buffalo.Simple
func (w SimpleWorker) PerformIn(j Job, d time.Duration) error {
	w.Increment()
	return w.Simple.PerformIn(j, d)
}
