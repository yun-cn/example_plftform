package noonde

import (
	"fmt"

	worker "github.com/yanshiyason/noonde_platform/worker_memory"
)

// JobQueue call methods on JobQueue to start some background work.
var JobQueue jobQueue

// jobQueue wraps a worker and defines possible commands
type jobQueue struct {
	Worker
}

// InstabaseFetchListingReviews .
func (q jobQueue) InstabaseFetchListingReviews(listingID string) {
	q.Worker.Perform(Job{
		Queue:   "default",
		Handler: "instabase_fetch_listing_reviews",
		Args: Args{
			"listingID": listingID,
		},
	})
}

func (q jobQueue) SpacemarketFetchListingReviews(listingID int) {
	q.Worker.Perform(Job{
		Queue:   "default",
		Handler: "spacemarket_fetch_listing_reviews",
		Args: Args{
			"listingID": listingID,
		},
	})
}

func init() {
	fmt.Println("INITIALIZING SIMPLE WORKER ...")
	JobQueue = jobQueue{worker.NewSimpleWorker()}
	JobQueue.Worker.Register("instabase_fetch_listing_reviews", func(args worker.Args) error {
		listingID := args["listingID"].(string)
		instabaseQueryReviews(listingID)
		return nil
	})

	JobQueue.Worker.Register("spacemarket_fetch_listing_reviews", func(args worker.Args) error {
		listingID := args["listingID"].(int)
		spacemarketQueryReviews(listingID)
		return nil
	})
}
