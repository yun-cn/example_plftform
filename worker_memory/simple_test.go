package worker_test

import (
	"testing"
	"time"

	worker "github.com/yanshiyason/noonde_platform/worker_memory"
)

func addJob(w *worker.SimpleWorker, handler string) {
	w.Perform(worker.Job{
		Queue:   "default",
		Handler: handler,
		Args:    worker.Args{},
	})
}

// TODO test if adding removing jobs concurrently messes up the counter..
func Test_SimpleWorker(t *testing.T) {
	w := worker.NewSimpleWorker()
	w.Register("test_handler", func(args worker.Args) error {
		time.Sleep(10 * time.Millisecond)
		return nil
	})

	addJob(w, "test_handler")
	addJob(w, "test_handler")
	addJob(w, "test_handler")

	count := w.JobsInProgress

	if *count != 3 {
		t.Errorf("Was expecting 3 job enqueued had %d", *count)
	}

	time.Sleep(20 * time.Millisecond)

	count = w.JobsInProgress

	if *count != 0 {
		t.Errorf("Was expecting 0 job enqueued had %d", *count)
	}
}
