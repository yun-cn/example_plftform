package noonde

import (
	buffalo "github.com/gobuffalo/buffalo/worker"
)

type (
	// Worker implement this interface to do background job processing.
	Worker = buffalo.Worker
	// Job alias
	Job = buffalo.Job
	// Handler alias
	Handler = buffalo.Handler
	// Args alias
	Args = buffalo.Args
)
